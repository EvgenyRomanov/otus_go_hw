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

	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/app/scheduler"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/pkg/rmq"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/scheduler_config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGTSTP)
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

	rmqInstance := rmq.NewRmq(
		config.Rmq.ConsumerTag,
		config.Rmq.URI,
		config.Rmq.Exchange.Name,
		config.Rmq.Exchange.Type,
		config.Rmq.Exchange.QueueName,
		config.Rmq.Exchange.BindingKey,
		config.Rmq.MaxInterval,
	)

	err := rmqInstance.Connect()
	if err != nil {
		logg.Error("%s", "cannot connect to AMQP server: "+err.Error())
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()

		_, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := rmqInstance.Shutdown(); err != nil {
			logg.Error("%s", "failed to shutdown RMQ server: "+err.Error())
		} else {
			logg.Info("RMQ server successfully terminated!")
		}
	}()

	scheduler := scheduler.New(
		logg,
		eventStorage,
		rmqInstance,
		config.Scheduler.RunFrequencyInterval,
		config.Scheduler.TimeForRemoveOldEvents,
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler.NotificationSender(ctx)
	}()
	wg.Wait()
}
