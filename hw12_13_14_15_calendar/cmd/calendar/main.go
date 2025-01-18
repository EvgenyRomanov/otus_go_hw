package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/calendar_config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config := NewConfig()
	logg := logger.New(config.Logger.Level, os.Stdout)

	var eventStorage storage.EventStorage

	if config.Storage.Driver == "postgres" {
		connectionString := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.DB.DBHost, config.DB.DBPort, config.DB.DBUsername, config.DB.DBPassword, config.DB.DBName,
		)

		eventStorage = sqlstorage.New(connectionString, config.Storage.MigrationsPath)
		err := eventStorage.Connect(ctx)
		if err != nil {
			logg.Error("%s", "cannot connect to DB server: "+err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
		defer eventStorage.Close()
	} else {
		eventStorage = memorystorage.New()
	}

	logg.Info("%s", fmt.Sprintf("successfully init %s storage", config.Storage.Driver))

	calendar := app.New(logg, eventStorage)

	httpServer := internalhttp.NewServer(config.HTTPServer.Host, config.HTTPServer.Port, logg, calendar)
	grpcServer := grpc.NewServer(config.GRPCServer.Host, config.GRPCServer.Port, logg, calendar)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			logg.Error("%s", "failed to stop http server: "+err.Error())
		}

		logg.Info("http-server successfully terminated!")

		grpcServer.Stop()
		logg.Info("grpc-server successfully terminated!")

		os.Exit(1)
	}()

	logg.Info("calendar is running...")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		if err := httpServer.Start(ctx); err != nil {
			logg.Error("%s", "failed to start http server: "+err.Error())
			wg.Done()
		}
	}()

	go func() {
		if err := grpcServer.Start(ctx); err != nil {
			logg.Error("%s", "failed to start grpc server: "+err.Error())
			wg.Done()
		}
	}()

	wg.Wait()
}
