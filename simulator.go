package main

import (
	"io/ioutil"
	"time"
)

type simulator struct {
	id string
	user
}

type stream struct {
	id     string
	result streamResult
}

type streamResult struct {
	status        int
	contentLength int64
	timeTaken     time.Duration
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

func click(http mainHttpClient, url string) streamResult {
	response, timeTaken, err := sendRequest(url)
	defer response.Body.Close()

	if err != nil {
		return streamResult{
			status:        0,
			contentLength: 0,
			timeTaken:     0,
		}
	}

	if response.ContentLength == -1 {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return streamResult{
				status:        response.StatusCode,
				contentLength: 0,
				timeTaken:     timeTaken,
			}
		}

		return streamResult{
			status:        response.StatusCode,
			contentLength: int64(len(body)),
			timeTaken:     timeTaken,
		}
	}

	return streamResult{
		status:        response.StatusCode,
		contentLength: response.ContentLength,
		timeTaken:     timeTaken,
	}
}

func newStream(id string, result streamResult) stream {
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
