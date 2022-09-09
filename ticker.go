package main

import "time"

type ticker struct {
	tick chan bool
}

func newTicker(interval int) ticker {
	tick := make(chan bool)
	ticker := ticker{tick: tick}

	go func(t chan bool) {
		for {
			m := time.Now().Unix() + int64(interval)

			for {
				if time.Now().Unix() > m {
					ticker.tick <- true
				}
			}
		}
	}(tick)

	return ticker
}