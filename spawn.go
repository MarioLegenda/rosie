package main

type spawnData struct {
	stream chan stream
	output chan status
	http   mainHttpClient
}

func newSpawnData() spawnData {
	return spawnData{
		stream: make(chan stream),
		output: make(chan status),
		http:   newHttp(),
	}
}

func spawn(args arguments, simulators []simulator, spawnData spawnData) chan stream {
	if !args.throttle {
		for _, s := range simulators {
			simulate(spawnData.http, s, spawnData.stream, spawnData.output)
		}
	}

	go func() {
		for s := range spawnData.output {
			for _, t := range simulators {
				if t.id == s.id {
					simulate(spawnData.http, t, spawnData.stream, spawnData.output)
				}
			}
		}
	}()

	return spawnData.stream
}
