package sort_pkg

import (
	"strconv"
	"strings"
	"unicode"
)

// sortRowByColumn represents slice of slices of strings
type sortRowByColumn struct {
	data        [][]string
	keyColumn   int
	sortNumeric bool
}

func (row sortRowByColumn) Len() int {
	return len(row.data)
}

func (row sortRowByColumn) Less(i, j int) bool {
	col := row.keyColumn
	if col > len(row.data[i])-1 || col > len(row.data[j]) {
		col = 0
	}

	if row.sortNumeric {
		n1 := strings.TrimFunc(row.data[i][col], func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		n2 := strings.TrimFunc(row.data[j][col], func(r rune) bool {
			return !unicode.IsNumber(r)
		})

		i1, err := strconv.Atoi(n1)
		if err != nil {
			return row.data[i][col] < row.data[j][col]
		}
		j1, err := strconv.Atoi(n2)
		if err != nil {
			return row.data[i][col] < row.data[j][col]
		}

		return i1 < j1
	}
	return row.data[i][col] < row.data[j][col]
}

func (row sortRowByColumn) Swap(i, j int) {
	row.data[i], row.data[j] = row.data[j], row.data[i]
}
