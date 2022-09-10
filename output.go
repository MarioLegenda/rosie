package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"time"
)

const DATE_LAYOUT = "2006-01-02T15:04:05Z"

type output struct {
	id              string
	name            string
	lastName        string
	totalRequestNum int
	success         int
	failed          int
}

type total struct {
	success     int
	failed      int
	total       int
	elapsedTime int64
}

func buildOutput(total total, table *tm.Table) {
	tm.Clear()

	fmt.Fprintf(table, "%d\t%d\t%d\t%d\n", total.success, total.failed, total.total, total.elapsedTime)

	tm.Print(table)
}

func createBaseTable() *tm.Table {
	base := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(base, "Success\tFailed\tTotal\tElapsed time\n")

	return base
}

func showTable() {
	tm.Flush()
}

func createOutputs(users []user) []*output {
	outputs := make([]*output, 0)
	for _, u := range users {
		o := &output{
			id:              u.UUID,
			name:            u.Name,
			lastName:        u.LastName,
			totalRequestNum: 0,
			success:         0,
			failed:          0,
		}

		outputs = append(outputs, o)
	}

	return outputs
}

func watchOutput(outputs []*output, st chan stream) {
	buildOutput(total{
		success:     0,
		failed:      0,
		total:       0,
		elapsedTime: 0,
	}, createBaseTable())
	showTable()
	start := time.Now().Unix()
	// calculate request fail/success
	go func(stream chan stream, outputs []*output) {
		internalOutputs := outputs
		ttl := total{
			success: 0,
			failed:  0,
			total:   0,
		}

		for s := range stream {
			for i, o := range internalOutputs {
				if s.id == o.id {
					o.totalRequestNum++

					if s.result {
						o.success++
					} else {
						o.failed++
					}

					internalOutputs[i] = o

					break
				}
			}

			sc, t, f := 0, 0, 0
			for _, o := range internalOutputs {
				t += o.totalRequestNum
				sc += o.success
				f += o.failed
			}

			ttl.total = t
			ttl.success = sc
			ttl.failed = f
			ttl.elapsedTime = time.Now().Unix() - start

			buildOutput(ttl, createBaseTable())
			showTable()
		}
	}(st, outputs)
}
