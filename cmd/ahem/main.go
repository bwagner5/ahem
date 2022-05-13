package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	delay := flag.Duration("delay", envOrDefault("delay", time.Second*30), "duration string to delay shutting down this application after a SIGTERM")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("I've been interrupted with a %s!\n", sig)
		fmt.Printf("Gracefully Shutting Down in %s\n", delay)
		signalReceived := time.Now().UTC()
		//todo: do a timer and a ticker in a select
		ticker := time.NewTicker(time.Second * 5)
		for _ = range ticker.C {
			durationUntilShutdown := *delay - time.Since(signalReceived)
			if durationUntilShutdown <= 0 {
				fmt.Println("👋 FINALLY SHUTTING DOWN! BYE!")
				done <- true
			}
			fmt.Printf("Shutting down in %s\n", durationUntilShutdown)
		}
	}()

	fmt.Printf("awaiting SIGTERM signal (pid=%d)\n", os.Getpid())
	<-done
	fmt.Println("exiting")
}

func envOrDefault(key string, def time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		dur, err := time.ParseDuration(val)
		if err != nil {
			return def
		}
		return dur
	}
	return def
}
