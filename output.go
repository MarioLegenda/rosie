package main

import (
	"fmt"
	tm "github.com/buger/goterm"
)

type output struct {
	id              string
	name            string
	lastName        string
	totalRequestNum int
	success         int
	failed          int
}

func buildOutput(output []output, table *tm.Table) {
	for i, o := range output {
		fmt.Fprintf(table, "%d\t%s\t%s\t%d\t%d\t%d\n", i+1, o.name, o.lastName, o.success, o.failed, o.totalRequestNum)
	}

	tm.Print(table)
}

func createBaseTable() *tm.Table {
	tm.Clear()

	base := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(base, "ID\tName\tLast name\tSuccess\tFailed\tR. num\n")

	return base
}

func showTable() {
	tm.Flush()
}

func createOutputs(users []user) []output {
	outputs := make([]output, 0)
	for _, u := range users {
		o := output{
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

func watchOutput(outputs []output, st chan stream) {
	buildOutput(outputs, createBaseTable())
	showTable()

	go func(stream chan stream, outputs []output) {
		internalOutputs := outputs
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

				}
			}

			buildOutput(internalOutputs, createBaseTable())
			showTable()
		}
	}(st, outputs)
}
