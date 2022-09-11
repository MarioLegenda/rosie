package main

import (
	"context"
	"fmt"
	pb "github.com/schollz/progressbar/v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type exit struct {
	stop  chan bool
	ctx   context.Context
	close func()
}

func newExit() exit {
	ctx, cancel := context.WithCancel(context.Background())

	return exit{
		stop: make(chan bool),
		ctx:  ctx,
		close: func() {
			cancel()
		},
	}
}

func initInterval(duration int) {
	go func() {
		t := newTicker(time.Duration(duration) * time.Second)

		for _ = range t.tick {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)

			break
		}
	}()
}

func initProgressBar(duration int) {
	go func() {
		bar := pb.Default(int64(duration))
		for i := 0; i < duration; i++ {
			bar.Add(1)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
}

func watchExit(exit exit, ttlCh chan total) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("")
	fmt.Println("Closing all simulators...")
	exit.close()

	ttl := <-ttlCh
	fmt.Println("Simulators closed. Collecting information...")
	fmt.Println("")

	fmt.Printf("Success: %d\n", ttl.success)
	fmt.Printf("Failed: %d\n", ttl.failed)
	fmt.Printf("Total: %d\n", ttl.total)
	fmt.Println("")

	go func() {
		exit.stop <- true
	}()
}
