package main

import (
	greppkg "grep/grep-pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrepCount(t *testing.T) {
	input := []string{
		"1", "1", "1", "1", "1", "6", "7", "8",
	}
	want := 5
	grepApp := greppkg.GrepApp{
		Pattern: "1",
		Count:   true,
	}
	count := grepApp.CountLines(input)
	if want != count {
		t.Fatalf(`Result = %d, want match for %d`, count, want)
	}
}

func TestGrep(t *testing.T) {
	testCases := []struct {
		desc string
		app  greppkg.GrepApp
		data []string
		want []string
	}{
		{
			desc: "normal",
			app: greppkg.GrepApp{
				Pattern: "4",
			},
			data: []string{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			want: []string{
				"4\n",
			},
		},
		{
			desc: "inverted",
			app: greppkg.GrepApp{
				Pattern: "4",
				Invert:  true,
			},
			data: []string{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			want: []string{
				"1\n", "2\n", "3\n", "5\n", "6\n", "7\n", "8\n",
			},
		},
		{
			desc: "before",
			app: greppkg.GrepApp{
				Pattern: "4",
				Before:  2,
			},
			data: []string{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			want: []string{
				"4\n", "Before match\n", "2\n", "3\n",
			},
		},
		{
			desc: "after",
			app: greppkg.GrepApp{
				Pattern: "4",
				After:   2,
			},
			data: []string{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			want: []string{
				"4\n", "After match\n", "5\n", "6\n",
			},
		},
		{
			desc: "context",
			app: greppkg.GrepApp{
				Pattern: "4",
				Context: 2,
			},
			data: []string{
				"1", "2", "3", "4", "5", "6", "7", "8",
			},
			want: []string{
				"4\n", "Context lines\n", "2\n", "3\n", "5\n", "6\n",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.app.GetLines(tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
