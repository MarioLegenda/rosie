package main

import (
	"time"
)

type ticker struct {
	tick chan bool
}

func newTicker(duration time.Duration) ticker {
	tick := make(chan bool)
	tckr := ticker{tick: tick}

	go func(t ticker) {
		for {
			tk := time.NewTicker(duration)

			for _ = range tk.C {
				t.tick <- true

				tk.Stop()

				close(t.tick)

				return
			}
		}
	}(tckr)

	return tckr
}
