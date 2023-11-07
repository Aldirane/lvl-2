package main

import (
	anagrampkg "anagram/anagram-pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		desc string
		data []string
		want map[string][]string
	}{
		{
			desc: "normal",
			data: []string{
				"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик", "листок",
			},
			want: map[string][]string{
				"листок": {"слиток", "столик"},
				"пятак":  {"пятка", "тяпка"},
			},
		},
		{
			desc: "none",
			data: []string{
				"пятак", "столик",
			},
			want: map[string][]string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			anagram := anagrampkg.Anagram{}
			got := anagram.BuildAnagram(&tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}
