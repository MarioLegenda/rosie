package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"log"
	"os"
	"time"
)

func main() {
	tm.Clear()

	run()
}

func coldStart(http http, urls []string) {
	for _, url := range urls {
		if value := click(http, url); !value {
			log.Fatal(fmt.Sprintf("URL %s did not return status code 200", url))
		}
	}
}

func run() {
	args, err := newArgs(os.Args[1:])

	if err != nil {
		log.Fatal(err.Error())
	}

	stop := make(chan bool)

	http := newHttp()

	fmt.Println("")
	fmt.Println("Initiating cold start...")
	coldStart(http, args.links)
	fmt.Println("Cold start finished. Sleeping for 10 seconds to give the server time to prepare for real testing...")
	time.Sleep(time.Second * 10)

	users := createUsers(args.users, args.links, args.intervalMin, args.intervalMax)
	outputs := createOutputs(users)
	simulators := createSimulator(users)

	stream := spawn(http, simulators)

	watchOutput(outputs, stream)

	<-stop
}
