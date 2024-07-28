package dp

import "strings"

func wordBreakV2(s string, wordDict []string) []string {
	wordSet := map[string]bool{}
	for _, w := range wordDict {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([][][]string, n)
	var sentences []string
	for _, words := range wordBreakV2Backtrack(s, wordSet, 0, &dp) {
		sentences = append(sentences, strings.Join(words, " "))
	}

	return sentences
}

func wordBreakV2Backtrack(s string, wordSet map[string]bool, i int, dp *[][][]string) [][]string {
	if (*dp)[i] != nil {
		return (*dp)[i]
	}

	var wordsList [][]string
	for j := i + 1; j < len(s); j++ {
		word := s[i:j]
		if wordSet[word] {
			for _, nextWords := range wordBreakV2Backtrack(s, wordSet, j, dp) {
				wordsList = append(wordsList, append([]string{word}, nextWords...))
			}
		}
	}

	word := s[i:]
	if wordSet[word] {
		wordsList = append(wordsList, []string{word})
	}
	(*dp)[i] = wordsList

	return wordsList
}
