package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"log"
	"os"
	"therebelsource/simulation/logger"
	"time"
)

func main() {
	tm.Clear()

	run()
}

func coldStart(urls []string) {
	for _, url := range urls {
		if value := click(url); !value {
			log.Fatal(fmt.Sprintf("URL %s did not return status code 200", url))
		}
	}
}

func run() {
	args, err := newArgs(os.Args[1:])
	/*
		urls := []string{
			"http://localhost:8080/",
			"http://localhost:8080/model/zs-ev",
			"http://localhost:8080/model/marvel-r",
			"http://localhost:8080/model/mg5",
			"http://localhost:8080/model/mg4",
		}
	*/
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.BuildLoggers()

	fmt.Println("")
	fmt.Println("Initiating cold start...")
	coldStart(args.links)
	fmt.Println("Cold start finished. Sleeping for 10 seconds to give the server time to prepare for real testing...")
	time.Sleep(time.Second * 10)

	stop := make(chan bool)

	users := createUsers(args.users, args.links, args.intervalMin, args.intervalMax)
	outputs := createOutputs(users)
	simulators := createSimulator(users)

	stream := spawn(simulators)

	watchOutput(outputs, stream)

	<-stop
}
