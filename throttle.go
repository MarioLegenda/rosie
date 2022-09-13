package main

import (
	"fmt"
	pb "github.com/schollz/progressbar/v3"
	"time"
)

func throttle(spawnData spawnData, simulators []simulator) {
	lc := newLoadCalc(len(simulators), 10)

	fmt.Println("Throttling requests in preparation to load testing...")
	fmt.Println("")

	bar := pb.Default(int64(lc.total + lc.remainder))

	curr := 0
	for i := 0; i < lc.total; i++ {
		t := newTicker(time.Duration(100) * time.Millisecond)

		for _ = range t.tick {
			for a := 0; a < lc.parts; a++ {
				s := simulators[curr]
				simulate(spawnData.http, s, spawnData.stream, spawnData.output)
				curr++
			}

			bar.Add(1)
		}
	}

	for i := 0; i < lc.remainder; i++ {
		s := simulators[curr]
		simulate(spawnData.http, s, spawnData.stream, spawnData.output)
		curr++

		bar.Add(lc.remainder)
	}

	bar.Close()

	fmt.Println("")
	fmt.Println("Throttling finished. Starting load testing.")
	fmt.Println("")
	time.Sleep(1 * time.Second)
}
