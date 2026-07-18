package tq_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tq "github.com/tsvsheet/go-tq"
)

// TestProgramConcurrentRuns verifies one Program value runs concurrently over
// different tables race-free (`make ci` runs the race detector): every stage
// kind that compiles expressions — where, derive, group — plus sort and
// distinct, sharing a single immutable Program.
func TestProgramConcurrentRuns(t *testing.T) {
	t.Parallel()
	program, err := tq.Parse(tq.Query(
		"where [n] > 0 | derive double = [n] * 2 | sort -double | distinct k | group k { total = sum([double]) }",
	))
	require.NoError(t, err)

	const runs = 16
	results := make([]string, runs)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := strconv.Itoa(i + 1)
			table := tq.Table{
				Header: []string{"k", "n"},
				Rows:   [][]string{{"a", n}, {"a", "0"}, {"b", n}},
			}
			out, err := program.Run(table, tq.Options{})
			assert.NoError(t, err)
			var got string
			for _, row := range out.Rows {
				got += row[0] + "=" + row[1] + ";"
			}
			results[i] = got
		}()
	}
	wg.Wait()
	for i, got := range results {
		double := strconv.Itoa((i + 1) * 2)
		assert.Equal(t, "a="+double+";b="+double+";", got, "run %d", i)
	}
}
