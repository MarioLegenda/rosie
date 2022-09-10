package main

func spawn(http http, simulators []simulator) chan stream {
	stream := make(chan stream)
	stat := make(chan status)

	for _, s := range simulators {
		simulate(http, s, stream, stat)
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
