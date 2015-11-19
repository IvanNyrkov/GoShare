package passphrase

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
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
	_, currentPath, _, _ := runtime.Caller(1)
	path := filepath.Dir(currentPath)

	nounsFilePath := filepath.Join(path, nounsFile)
	nounWords = mustInitWordsFromFile(nounsFilePath)

	adjectivesFilePath := filepath.Join(path, adjectivesFile)
	adjectiveWords = mustInitWordsFromFile(adjectivesFilePath)
}

// Initialization of words array from file
func mustInitWordsFromFile(filePath string) (words []string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
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
