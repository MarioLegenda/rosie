package main

func spawn(simulators []simulator) chan stream {
	stream := make(chan stream)
	stat := make(chan status)

	for _, s := range simulators {
		simulate(s, stream, stat)
	}

	go func() {
		for s := range stat {
			for _, t := range simulators {
				if t.id == s.id {
					simulate(t, stream, stat)
				}
			}
		}
	}()

	return stream
}
