package main

import (
	"fmt"
	"log"
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
		if value := click(http, url); !value {
			log.Fatal(fmt.Sprintf("URL %s did not return status code 200", url))
		}
	}
	fmt.Println("Cold start finished. Sleeping for 10 seconds to give the server time to prepare for real testing...")
}

func run() {
	args, err := newArgs(os.Args[1:])

	if err != nil {
		log.Fatal(err.Error())
	}

	http := newHttp()

	coldStart(http, args.links)
	time.Sleep(time.Second * 10)

	stop := newExit()

	users := createUsers(args)
	simulators := createSimulator(users)
	data := newSpawnData()

	if args.throttle {
		throttle(data, simulators)
	} else {
		fmt.Println("Running load requests now...")
		fmt.Println("")
	}

	initInterval(args.duration)
	initProgressBar(args.duration)
	watchExit(stop, watchOutput(createOutputs(users), spawn(args, simulators, data), stop))

	<-stop.stop
}
