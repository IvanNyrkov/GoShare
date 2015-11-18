package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

var (
	nounsFile      = "nouns.txt"
	adjectivesFile = "adjectives.txt"
	nounWords      []string
	adjectiveWords []string
)

// Initialization of word arrays
func init() {
	nounWords = mustInitWordsFromFile(nounsFile)
	adjectiveWords = mustInitWordsFromFile(adjectivesFile)
}

// Initialization of words array from file
func mustInitWordsFromFile(filename string) (words []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return
}

// GetRandomSentence returns random generated sentence of adjective + noun
func GetRandomSentence(separator string) string {
	nounRand := rand.Intn(len(nounWords))
	adjectiveRand := rand.Intn(len(adjectiveWords))
	result := []string{
		adjectiveWords[adjectiveRand],
		nounWords[nounRand],
	}
	return strings.Join(result, separator)
}
