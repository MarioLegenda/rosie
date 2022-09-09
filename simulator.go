package main

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

func click(url string) bool {
	return sendRequest(url)
}

func newStream(id string, result bool) stream {
	return stream{
		id:     id,
		result: result,
	}
}

func simulate(s simulator, st chan stream, stat chan status) {
	go func(s simulator, str chan stream, stat chan status) {
		user := s.user
		urls := user.urls

		ticker := newTicker(random(user.Min, user.Max))
		nextClick := 0

		for _ = range ticker.tick {
			if nextClick < len(urls) {
				st <- newStream(user.UUID, click(urls[nextClick]))
				nextClick++
			} else {
				stat <- status{id: s.id}
				return
			}
		}
	}(s, st, stat)
}

func createSimulator(users []user) []simulator {
	simulators := make([]simulator, 0)
	for _, u := range users {
		simulators = append(simulators, newSimulator(u))
	}

	return simulators
}
