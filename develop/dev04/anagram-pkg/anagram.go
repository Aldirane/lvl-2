package anagram_pkg

import (
	"sort"
	"strings"
)

type Anagram struct {
	data map[string][]string
}

func (anagramApp *Anagram) BuildAnagram(input *[]string) map[string][]string {
	uniqueSet := anagramApp.uniqueWords(*input)
	anagramApp.data = anagramApp.makeAnagramMap(uniqueSet)
	return anagramApp.data
}

func (anagramApp *Anagram) uniqueWords(input []string) []string {
	uniqueSet := make([]string, 0)
	uniqueMap := make(map[string]bool)
	for _, word := range input {
		word = strings.ToLower(word)
		if _, ok := uniqueMap[word]; !ok {
			uniqueMap[word] = true
			uniqueSet = append(uniqueSet, word)
		}
	}
	return uniqueSet
}

func (anagramApp *Anagram) makeAnagramMap(uniqueSet []string) map[string][]string {
	anagramMap := make(map[string][]string)
	anagrams := make([]string, 0)
	idxInMap := make(map[int]bool)
	for idx1, word1 := range uniqueSet {
		idxInMap[idx1] = true
		for idx2, word2 := range uniqueSet {
			if !idxInMap[idx2] && compareRunes(word1, word2) {
				anagrams = append(anagrams, word2)
				idxInMap[idx2] = true
			}
		}
		if len(anagrams) > 0 {
			sort.SliceStable(anagrams, func(i int, j int) bool {
				return anagrams[i] < anagrams[j]
			})
			anagramMap[word1] = anagrams[:]
			anagrams = make([]string, 0)
		}
	}
	return anagramMap
}

func compareRunes(word1, word2 string) bool {
	charMap1 := make(map[string]bool)
	charMap2 := make(map[string]bool)
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1); i++ {
		charMap1[string(word1[i])] = true
	}
	for i := 0; i < len(word1); i++ {
		if _, ok := charMap1[string(word2[i])]; ok {
			charMap2[string(word1[i])] = true
		} else {
			return false
		}
	}
	if len(charMap1) != len(charMap2) {
		return false
	}
	return true
}
