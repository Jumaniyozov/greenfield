// Package translation provides translation functionalities for the application.
package translation

import "strings"

// StaticService has data that does not change.
type StaticService struct{}

// NewStaticService creates new instance of a service that uses static data.
func NewStaticService() *StaticService {
	return &StaticService{}
}

// Translate a given word to a the passed in language.
func (s *StaticService) Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}

	switch strings.ToLower(language) {
	case "english":
		return word
	case "german":
		return "hallo"
	case "finnish":
		return "hei"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
