package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/app/sender"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/pkg/rmq"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/sender_config.toml", "Path to configuration file")
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

	sender := sender.New(logg, rmqInstance, config.Sender.Threads)

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := sender.Consume(ctx)
		if err != nil {
			logg.Error("cannot init consumer for AMQP server: %s", err.Error())
		}
	}()
	wg.Wait()
}
