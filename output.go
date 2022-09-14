package main

type output struct {
	id              string
	name            string
	lastName        string
	totalRequestNum int
	success         int
	failed          int
}

func watchOutput(st chan stream, ext exit) chan []streamResult {
	stop := make(chan []streamResult)
	// calculate request fail/success
	go func(stream chan stream) {
		streamResults := make([]streamResult, 0)

		for {
			select {
			case <-ext.ctx.Done():
				stop <- streamResults

				return
			case s := <-st:
				streamResults = append(streamResults, s.result)
			}
		}
	}(st)

	return stop
}
