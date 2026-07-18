package tq_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tq "github.com/tsvsheet/go-tq"
)

// corpusFiles lists the .tq files under testdata/corpus/<kind> — the corpus
// copied verbatim from the grammar repo (tsvsheet/tq/testdata).
func corpusFiles(t *testing.T, kind string) []string {
	t.Helper()
	files, err := filepath.Glob(filepath.Join("testdata", "corpus", kind, "*.tq"))
	require.NoError(t, err)
	require.NotEmpty(t, files, "corpus %s must not be empty", kind)
	return files
}

// TestCorpusValidParses asserts every valid corpus program parses (R1) — no
// partial programs, no errors.
func TestCorpusValidParses(t *testing.T) {
	t.Parallel()
	for _, file := range corpusFiles(t, "valid") {
		t.Run(filepath.Base(file), func(t *testing.T) {
			t.Parallel()
			src, err := os.ReadFile(file)
			require.NoError(t, err)
			_, err = tq.Parse(tq.Query(src))
			assert.NoError(t, err)
		})
	}
}

// TestCorpusInvalidFailsSyntax asserts every invalid corpus program fails
// Parse with ErrSyntax (R1), matchable with errors.Is.
func TestCorpusInvalidFailsSyntax(t *testing.T) {
	t.Parallel()
	for _, file := range corpusFiles(t, "invalid") {
		t.Run(filepath.Base(file), func(t *testing.T) {
			t.Parallel()
			src, err := os.ReadFile(file)
			require.NoError(t, err)
			_, err = tq.Parse(tq.Query(src))
			require.Error(t, err)
			assert.True(t, errors.Is(err, tq.ErrSyntax), "want ErrSyntax, got %v", err)
		})
	}
}

// TestCorpusA1ParsesPlanRejects asserts the raw-A1 corpus program parses but
// fails PLAN with ErrCellRef: tq addresses columns, never cells (ADR 0002).
func TestCorpusA1ParsesPlanRejects(t *testing.T) {
	t.Parallel()
	src, err := os.ReadFile(filepath.Join("testdata", "corpus", "valid", "14-a1-parses-plan-rejects.tq"))
	require.NoError(t, err)
	program, err := tq.Parse(tq.Query(src))
	require.NoError(t, err)

	table, err := tq.ReadTable(strings.NewReader(repos), tq.Options{})
	require.NoError(t, err)
	_, err = program.Run(table, tq.Options{})
	require.Error(t, err)
	assert.True(t, errors.Is(err, tq.ErrCellRef), "want ErrCellRef, got %v", err)
}
