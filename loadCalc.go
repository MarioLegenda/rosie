package main

type loadCalc struct {
	total     int
	remainder int
	parts     int
}

func newLoadCalc(num int, parts int) loadCalc {
	return loadCalc{
		total:     num / parts,
		remainder: num % parts,
		parts:     parts,
	}
}
