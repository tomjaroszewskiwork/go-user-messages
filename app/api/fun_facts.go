package api

// Fun facts about a users message
type FunFacts struct {

	// Does the message have exciting content
	Exciting bool `json:"exciting"`

	// Is the message an paldindrome
	Palindrome bool `json:"palindrome"`

	// Does the message have sad content
	Sad bool `json:"sad"`
}
