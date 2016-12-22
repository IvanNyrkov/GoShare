package test

import (
	"testing"

	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
	"github.com/stretchr/testify/require"
)

// TestMustNotBeAbleToSetupNotExistingFiles we must not be able to setup the module with not existing files
func TestMustNotBeAbleToSetupNotExistingFiles(t *testing.T) {
	_, err := sentence.NewModule(sentence.ModuleConfig{
		NounsFile:      "missing.file",
		AdjectivesFile: "another.missing.file",
	})
	require.NotNil(t, err)
}

// TestMustBeAbleToSetupEmptyFiles we must be able to setup the module with empty files
func TestMustBeAbleToSetupEmptyFiles(t *testing.T) {
	_, err := sentence.NewModule(sentence.ModuleConfig{
		NounsFile:      "empty_file.txt",
		AdjectivesFile: "empty_file.txt",
	})
	require.Nil(t, err)
}

// TestMustReceiveEmptyStringIfFilesAreEmpty we must receive empty string if file is empty
func TestMustReceiveEmptyStringIfFilesAreEmpty(t *testing.T) {
	m, err := sentence.NewModule(sentence.ModuleConfig{
		NounsFile:      "empty_file.txt",
		AdjectivesFile: "empty_file.txt",
	})
	require.Nil(t, err)
	s := m.GetService().RandomSentence(".")
	require.Empty(t, s)
}

// TestMustBeAbleToSetupCorrectFiles we must be able to setup module with correct files
func TestMustBeAbleToSetupCorrectFiles(t *testing.T) {
	_, err := sentence.NewModule(sentence.ModuleConfig{
		NounsFile:      "one_word_file.txt",
		AdjectivesFile: "one_word_file.txt",
	})
	require.Nil(t, err)
}

// TestSeparator tests that we can set any separator for returned sentence
func TestSeparator(t *testing.T) {
	m, err := sentence.NewModule(sentence.ModuleConfig{
		NounsFile:      "one_word_file.txt",
		AdjectivesFile: "one_word_file.txt",
	})
	require.Nil(t, err)
	service := m.GetService()
	for _, s := range []string{".", "/", "|", "-", "_", "+", ""} {
		require.Equal(t, "test"+s+"test", service.RandomSentence(s))
	}
}

// TestMustNotBeAbleToUseEmptyConfig we must not be able to use empty config
func TestMustNotBeAbleToUseDefaultFiles(t *testing.T) {
	_, err := sentence.NewModule(sentence.ModuleConfig{})
	require.NotNil(t, err)
}

// TestMustBeAbleToUseDefaultFiles we must be able to not specify any files
func TestMustBeAbleToUseDefaultFiles(t *testing.T) {
	_, err := sentence.NewModule(sentence.DefaultConfig)
	require.Nil(t, err)
}

// TestSentenceMustBeRandom we must receive random result
func TestSentenceMustBeRandom(t *testing.T) {
	m, err := sentence.NewModule(sentence.DefaultConfig)
	require.Nil(t, err)
	service := m.GetService()
	// Try to get different result for 'limit' tries
	// TODO: It's random, so it can be properly tested
	firstResult := service.RandomSentence("")
	currentTry := 1
	limit := 50
	for firstResult == service.RandomSentence("") && currentTry != limit {
		currentTry++
	}
	require.NotEqual(t, limit, currentTry)
}
