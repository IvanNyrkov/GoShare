package sentence

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

// Service is an interface that defines actions for generation of random sentences
type Service interface {
	GetRandomSentence(sep string) string
}
type serviceImpl struct {
	adjectives []string
	nouns      []string
}

// NewService creates new service for generation random sentences
func NewService(adjFilePath, nounsFilePath string) (*serviceImpl, error) {
	adjectives, err := readWordsFromFile(adjFilePath)
	if err != nil {
		return nil, err
	}
	nouns, err := readWordsFromFile(nounsFilePath)
	if err != nil {
		return nil, err
	}
	return &serviceImpl{
		adjectives: adjectives,
		nouns:      nouns,
	}, nil
}

// GetRandomSentence returns random sentence (noun + adj) divided by the specified separator
func (s *serviceImpl) GetRandomSentence(sep string) string {
	adjAmount := len(s.adjectives)
	nounAmount := len(s.nouns)
	if adjAmount == 0 || nounAmount == 0 {
		return ""
	}
	adjectiveRand := rand.Intn(len(s.adjectives))
	nounRand := rand.Intn(len(s.nouns))
	result := []string{
		s.adjectives[adjectiveRand],
		s.nouns[nounRand],
	}
	return strings.Join(result, sep)
}

// readWordsFromFile parses file lines into string slice
func readWordsFromFile(filePath string) ([]string, error) {
	// Open file
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Read file line by line
	words := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}
