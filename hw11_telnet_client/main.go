package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/pflag"
)

var (
	timeout    time.Duration
	host, port string
)

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?

	pflag.DurationVarP(&timeout, "timeout", "t", time.Second*10, "connection timeout")
	pflag.Parse()

	if len(pflag.Args()) < 2 {
		log.Fatal("incorrect number of parameters")
	}

	host = pflag.Arg(0)
	port = pflag.Arg(1)

	tc := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)

	if err := StartTelnetClient(tc); err != nil {
		log.Fatal(err)
	}
}

func StartTelnetClient(conn TelnetClient) error {
	log.Println("Start connection")

	if err := conn.Connect(); err != nil {
		return fmt.Errorf("connection error: %w", err)
	}

	defer func(conn TelnetClient) {
		_ = conn.Close()
	}(conn)

	log.Println("Connection success")

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, syscall.SIGINT)
	defer close(osSignalChan)

	doneChan := make(chan struct{})

	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := conn.Send(); err != nil {
			log.Println("Send error: ", err)
			doneChan <- struct{}{}
			return
		}
	}()

	go func() {
		defer wg.Done()

		if err := conn.Receive(); err != nil {
			log.Println("Receive error: ", err)
			doneChan <- struct{}{}
			return
		}
	}()

	for {
		select {
		case <-osSignalChan:
			return nil
		case <-doneChan:
			return nil
		}
	}
}
