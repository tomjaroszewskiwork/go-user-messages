package api

import (
	"strings"
)

// FunFacts about a users message
type FunFacts struct {

	// Does the message have exciting content
	Exciting bool `json:"exciting"`

	// Is the message an paldindrome
	Palindrome bool `json:"palindrome"`

	// Does the message have sad content
	Sad bool `json:"sad"`
}

// NewFuncFacts builds new fun facts for the message
func NewFuncFacts(message string) *FunFacts {
	newFacts := FunFacts{
		Exciting:   isExciting(message),
		Palindrome: isPallindrome(message),
		Sad:        isSad(message),
	}
	return &newFacts
}

func isPallindrome(message string) bool {
	n := len(message)
	for i := 0; i < n/2; i++ {
		if message[i] != message[n-i-1] {
			return false
		}
	}
	return true
}

func isExciting(message string) bool {
	return strings.Contains(message, "!") || strings.ToUpper(message) == message
}

func isSad(message string) bool {
	return strings.Contains(message, ":(") || strings.Contains(message, ":-(")
}
