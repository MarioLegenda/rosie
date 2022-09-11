package main

type loadCalc struct {
	total     int
	remainder int
	parts     int
}

func newLoadCalc(num int, parts int) loadCalc {
	t := num / parts
	r := num % parts
	return loadCalc{
		total:     t,
		remainder: r,
		parts:     parts,
	}
}
