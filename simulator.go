package main

import (
	"time"
)

type simulator struct {
	id string
	user
}

type stream struct {
	id     string
	result bool
}

type status struct {
	id string
}

func newSimulator(user user) simulator {
	return simulator{
		user: user,
		id:   user.UUID,
	}
}

func click(http mainHttpClient, url string) bool {
	response, _ := sendRequest(http, url)

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return false
	}

	return true
}

func newStream(id string, result bool) stream {
	return stream{
		id:     id,
		result: result,
	}
}

func simulate(http mainHttpClient, s simulator, st chan stream, stat chan status) {
	go func(s simulator, str chan stream, stat chan status) {
		user := s.user
		urls := user.urls

		for _, u := range urls {
			ticker := newTicker(time.Duration(random(user.Min, user.Max)) * time.Second)

			for _ = range ticker.tick {
				st <- newStream(user.UUID, click(http, u))

				break
			}
		}

		stat <- status{id: s.id}
	}(s, st, stat)
}

func createSimulator(users []user) []simulator {
	simulators := make([]simulator, 0)
	for _, u := range users {
		simulators = append(simulators, newSimulator(u))
	}

	return simulators
}
