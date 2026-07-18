package tq_test

import (
	"os"
	"strings"

	tq "github.com/tsvsheet/go-tq"
)

// Example runs the canonical pipeline from SPECIFICATION §3: filter, derive,
// project, sort, and limit — TSV in, TSV out.
func Example() {
	input := "name\tstars\tforks\n" +
		"alpha\t1500\t3\n" +
		"beta\t900\t2\n" +
		"gamma\t2000\t8\n"

	program, err := tq.Parse(
		"where [stars] > 1000 | derive ratio = round([stars] / [forks], 2) | select name, ratio | sort -ratio",
	)
	if err != nil {
		panic(err)
	}
	table, err := tq.ReadTable(strings.NewReader(input), tq.Options{})
	if err != nil {
		panic(err)
	}
	out, err := program.Run(table, tq.Options{})
	if err != nil {
		panic(err)
	}
	if err := tq.WriteTable(os.Stdout, out); err != nil {
		panic(err)
	}
	// Output:
	// name	ratio
	// alpha	500
	// gamma	250
}
