package exec

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Key is one resolved sort key: the column position and whether `-` reversed
// the key's entire order.
type Key struct {
	At           Position
	IsDescending bool
}

// Keys is a sort's resolved multi-key list.
type Keys []Key

// stableSort orders rows in place (over the caller's fresh copy) by the §5
// total order, key by key.
func stableSort(rows Rows, keys Keys) {
	sort.SliceStable(rows, func(i, j int) bool {
		return rowLess(Row(rows[i]), Row(rows[j]), keys)
	})
}

// rowLess compares two rows key by key; equal keys fall through so the stable
// sort preserves their prior relative order.
func rowLess(a, b Row, keys Keys) bool {
	for _, k := range keys {
		c := compareCells(cellText(cellAt(a, k.At)), cellText(cellAt(b, k.At)))
		if k.IsDescending {
			c = -c
		}
		if c != 0 {
			return c < 0
		}
	}
	return false
}

// compareCells is the per-key total order (ADR 0003): cells that coerce to
// numbers order before cells that do not; numeric cells order numerically
// among themselves; non-numeric cells order by byte-wise comparison of raw
// text among themselves. Numeric coercion is exactly the engine's literal
// rule — strconv.ParseFloat succeeds — so tq and a sheet agree on what a
// number is.
func compareCells(a, b cellText) int {
	an, aok := numericCell(a)
	bn, bok := numericCell(b)
	switch {
	case aok && bok:
		return compareFloats(an, bn)
	case aok:
		return -1
	case bok:
		return 1
	default:
		return strings.Compare(string(a), string(b))
	}
}

// numericKey is a cell's numeric coercion, as a sort key.
type numericKey float64

// numericCell coerces a cell to a number by the engine's literal rule.
func numericCell(s cellText) (numericKey, bool) {
	n, err := strconv.ParseFloat(string(s), 64)
	return numericKey(n), err == nil
}

// compareFloats orders two numeric keys, keeping totality when a cell parses
// to NaN: NaN equals NaN and orders after every other number.
func compareFloats(a, b numericKey) int {
	switch {
	case math.IsNaN(float64(a)) && math.IsNaN(float64(b)):
		return 0
	case math.IsNaN(float64(a)):
		return 1
	case math.IsNaN(float64(b)):
		return -1
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
