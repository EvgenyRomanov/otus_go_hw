package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const TopCount = 10

var (
	PunctuationMarks = ".,-!?:;"
	EmptyString      = ""
)

type Word struct {
	Word string
	Cnt  int
}

func Top10(text string, taskWithAsteriskIsCompleted bool) []string {
	textSplit := strings.Fields(text)
	wordMap := make(map[string]int)
	wordSlice := make([]Word, 0, TopCount)
	result := make([]string, 0, TopCount)

	for _, word := range textSplit {
		if taskWithAsteriskIsCompleted {
			word = strings.ToLower(word)
			word = strings.TrimLeftFunc(word, TrimCallback)
			word = strings.TrimRightFunc(word, TrimCallback)

			if word == EmptyString {
				continue
			}
		}

		_, ok := wordMap[word]

		if ok {
			wordMap[word]++
		} else {
			wordMap[word] = 1
		}
	}

	for word, cnt := range wordMap {
		wordSlice = append(wordSlice, Word{Word: word, Cnt: cnt})
	}

	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].Cnt == wordSlice[j].Cnt {
			return wordSlice[i].Word < wordSlice[j].Word
		}
		return wordSlice[i].Cnt > wordSlice[j].Cnt
	})

	topBorder := TopCount
	lenWordSlice := len(wordSlice)

	if lenWordSlice < TopCount {
		topBorder = lenWordSlice
	}

	for _, word := range wordSlice[:topBorder] {
		result = append(result, word.Word)
	}

	return result
}

func TrimCallback(r rune) bool {
	return strings.ContainsRune(PunctuationMarks, r)
}
