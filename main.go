package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

func main() {
	run()
}

func coldStart(http mainHttpClient, urls []string) {
	fmt.Println("")
	fmt.Println("Initiating cold start...")
	for _, url := range urls {
		result := click(http, url)

		if result.status < 200 || result.status > 299 {
			c := color.New(color.FgHiYellow).Add(color.Bold)
			c.Println(fmt.Sprintf("WARNING: URL %s returned status code %d.", url, result.status))
		}
	}
	fmt.Println("Cold start finished. Sleeping for 10 seconds to give the server time to prepare for real testing...")
}

func run() {
	args, ok := newArgs(os.Args[1:])

	if !ok {
		return
	}

	http := newHttp()

	coldStart(http, args.links)
	time.Sleep(time.Second * 10)

	stop := newExit()

	users := createUsers(args)
	simulators := createSimulator(users)
	data := newSpawnData()

	if args.throttle {
		fmt.Println("Throttling requests in preparation to load testing...")
		fmt.Println("")
	}

	if args.throttle {
		throttle(data, simulators)
	} else {
		fmt.Println("Running load requests now...")
		fmt.Println("")
	}

	initInterval(args.duration)
	initProgressBar(args.duration)
	watchExit(stop, watchOutput(spawn(args, simulators, data), stop))

	<-stop.stop
}
