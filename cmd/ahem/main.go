package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	delay := flag.Duration("delay", envOrDefault("delay", time.Second*25), "duration string to delay shutting down this application after a SIGTERM")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		<-sigs
		signalReceived := time.Now().UTC()
		log.Printf("I've been interrupted with a SIGTERM!")
		log.Printf("Gracefully Shutting Down in %d seconds", int(delay.Seconds()))

		timer := time.NewTimer(*delay)
		ticker := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-ticker.C:
				durationUntilShutdown := *delay - time.Since(signalReceived)
				if durationUntilShutdown > 0 {
					log.Printf("Shutting down in %d seconds", int(durationUntilShutdown.Seconds()))
				}
			case <-timer.C:
				log.Println("👋 FINALLY SHUTTING DOWN!")
				close(done)
				return
			}
		}
	}()
	log.Printf("awaiting SIGTERM signal (pid=%d)\n", os.Getpid())
	<-done
	log.Println("Bye!")
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
