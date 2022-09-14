package main

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	pb "github.com/schollz/progressbar/v3"
	"os"
	"os/signal"
	"strconv"
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

		bar.Close()
	}()
}

func watchExit(exit exit, ttlCh chan []streamResult) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	fmt.Println("")
	fmt.Println("Closing all simulators...")
	exit.close()

	streamResults := <-ttlCh
	fmt.Println("Simulators closed. Collecting information...")
	fmt.Println("")

	statusMap := make(map[string]int)
	var totalContentLen int64
	var failedRequests int64

	for _, s := range streamResults {
		if s.status == 0 {
			failedRequests++

			continue
		}

		status := strconv.Itoa(s.status)

		if _, ok := statusMap[status]; !ok {
			statusMap[status] = 0
		}

		statusMap[status]++
		totalContentLen += s.contentLength
	}

	for k, v := range statusMap {
		fmt.Println(fmt.Sprintf("Status code %s: %d", k, v))
	}
	fmt.Printf("Failed requests: %d\n", failedRequests)

	fmt.Println("")
	fmt.Printf("Total requests: %d\n", len(streamResults))
	fmt.Println(fmt.Sprintf("Total content length: %d bytes", totalContentLen))
	fmt.Println("")

	if len(streamResults) > 0 {
		min := streamResults[0].timeTaken
		for _, n := range streamResults {
			if n.status != 0 && n.timeTaken < min {
				min = n.timeTaken
			}
		}

		var max time.Duration
		for _, n := range streamResults {
			if n.status != 0 && n.timeTaken > max {
				max = n.timeTaken
			}
		}

		fmt.Printf("Fastest request: %.3vms\n", min)
		fmt.Printf("Slowest request: %.3vms\n", max)
		fmt.Println("")
	} else {
		warn := color.New(color.FgYellow).Add(color.Bold)
		warn.Println("WARNING: No information was able to be collected. This might be because the program has been interrupted before any requests could be sent.")
		fmt.Println("")
	}

	go func() {
		exit.stop <- true
	}()
}
