package main

import (
	"time"
)

type ticker struct {
	tick chan bool
}

func newTicker(interval int) ticker {
	tick := make(chan bool)
	tckr := ticker{tick: tick}

	go func(t ticker) {
		for {
			tk := time.NewTicker(time.Duration(interval) * time.Second)

			for _ = range tk.C {
				t.tick <- true

				tk.Stop()

				return
			}
		}
	}(tckr)

	return tckr
}
