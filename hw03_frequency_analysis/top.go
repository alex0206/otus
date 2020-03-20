package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"bufio"
	"log"
	"sort"
	"strings"
)

type word struct {
	name string
	cnt  int
}

func Top10(text string) []string {
	wordCntMap, err := wordCntAnalize(text)

	if err != nil {
		log.Fatalln(err.Error())
	}

	words := sortWords(wordCntMap)

	topWordSize := 10

	if wordsLen := len(words); wordsLen < topWordSize {
		topWordSize = wordsLen
	}

	topWords := make([]string, 0, topWordSize)

	for i := 0; i < topWordSize; i++ {
		topWords = append(topWords, words[i].name)
	}

	return topWords
}

func sortWords(wordCntMap map[string]int) []word {
	words := make([]word, 0, len(wordCntMap))

	for k, v := range wordCntMap {
		words = append(words, word{name: k, cnt: v})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].cnt > words[j].cnt })

	return words
}

func wordCntAnalize(text string) (map[string]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	wordCnt := make(map[string]int)

	for scanner.Scan() {
		wordCnt[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCnt, nil
}
