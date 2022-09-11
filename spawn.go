package main

import (
	"fmt"
	"time"
)

func spawn(args arguments, http http, simulators []simulator) chan stream {
	stream := make(chan stream)
	stat := make(chan status)

	if args.throttle {
		lc := newLoadCalc(len(simulators), 10)

		fmt.Println("Throttling requests in preparation to load testing...")
		curr := 0
		for i := 0; i < lc.total; i++ {
			t := newTicker(time.Duration(100) * time.Millisecond)

			for _ = range t.tick {
				for a := 0; a < lc.parts; a++ {
					s := simulators[curr]
					simulate(http, s, stream, stat)
					curr++
				}
			}
		}

		for i := 0; i < lc.remainder; i++ {
			s := simulators[curr]
			simulate(http, s, stream, stat)
			curr++
		}

		fmt.Println("Throttling finished. Starting load testing.")
		time.Sleep(1 * time.Second)
	} else {
		for _, s := range simulators {
			simulate(http, s, stream, stat)
		}
	}

	go func() {
		for s := range stat {
			for _, t := range simulators {
				if t.id == s.id {
					simulate(http, t, stream, stat)
				}
			}
		}
	}()

	return stream
}
