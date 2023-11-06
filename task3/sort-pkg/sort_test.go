package sort_pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		app  SortModule
		data []string
		want []string
	}{
		{
			desc: "normal sort",
			app:  SortModule{},
			data: []string{
				"first string",
				"second string",
				"third string",
				"fourth string",
				"fifth string",
			},
			want: []string{
				"fifth string",
				"first string",
				"fourth string",
				"second string",
				"third string",
			},
		},
		{
			desc: "not numeric sort without sortNumeric flag",
			app:  SortModule{},
			data: []string{
				"1",
				"2",
				"3",
				"5",
				"11",
				"22",
				"33",
				"55",
			},
			want: []string{
				"1",
				"11",
				"2",
				"22",
				"3",
				"33",
				"5",
				"55",
			},
		},
		{
			desc: "not numeric sort reverse order",
			app: SortModule{
				reverseSort: true,
			},
			data: []string{
				"1",
				"11",
				"2",
				"22",
				"3",
				"33",
				"5",
				"55",
			},
			want: []string{
				"55",
				"5",
				"33",
				"3",
				"22",
				"2",
				"11",
				"1",
			},
		},
		{
			desc: "delete duplicate",
			app: SortModule{
				removeDuplicates: true,
			},
			data: []string{
				"1",
				"1",
				"11",
				"11",
				"2",
				"2",
				"33",
				"33",
				"4",
				"4",
				"5",
			},
			want: []string{
				"1",
				"11",
				"2",
				"33",
				"4",
				"5",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortLines(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumns(t *testing.T) {
	testCases := []struct {
		desc string
		app  SortModule
		data []string
		want []string
	}{
		{
			desc: "sort by 2nd column",
			app: SortModule{
				keyColumn: 2,
			},
			data: []string{
				"string c",
				"string d",
				"string e",
				"string b",
				"string a",
			},
			want: []string{
				"string a",
				"string b",
				"string c",
				"string d",
				"string e",
			},
		},
		{
			desc: "sort by column out of range",
			app: SortModule{
				keyColumn: 10,
			},
			data: []string{
				"a string",
				"e string",
				"d string",
				"c string",
				"b string",
			},
			want: []string{
				"a string",
				"b string",
				"c string",
				"d string",
				"e string",
			},
		},
		{
			desc: "sort numbers numeric order",
			app: SortModule{
				keyColumn:   1,
				sortNumeric: true,
			},
			data: []string{
				"1",
				"11",
				"2",
				"22",
				"3",
				"33",
				"4",
			},
			want: []string{
				"1",
				"2",
				"3",
				"4",
				"11",
				"22",
				"33",
			},
		},
		{
			desc: "sort by 2nd column, in reverse",
			app: SortModule{
				keyColumn:   2,
				reverseSort: true,
			},
			data: []string{
				"string a",
				"string b",
				"string c",
				"string d",
				"string e",
			},
			want: []string{
				"string e",
				"string d",
				"string c",
				"string b",
				"string a",
			},
		},
		{
			desc: "delete duplicate",
			app: SortModule{
				keyColumn:        1,
				removeDuplicates: true,
			},
			data: []string{
				"1",
				"1",
				"2",
				"2",
				"3",
				"3",
				"3",
				"4",
				"4",
				"5",
				"5",
			},
			want: []string{
				"1",
				"2",
				"3",
				"4",
				"5",
			},
		},
		{
			desc: "numeric sort, but column starts with letter",
			app: SortModule{
				keyColumn:   1,
				sortNumeric: true,
			},
			data: []string{
				"a1a",
				"b5b",
				"c2c",
				"d3d",
				"e4e",
			},
			want: []string{
				"a1a",
				"c2c",
				"d3d",
				"e4e",
				"b5b",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.sortColumns(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
