package rest

import (
	"encoding/json"
	"github.com/jumaniyozov/greenfield/translation"
	"net/http"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translationWord := translation.Translate(word, language)
	if translationWord == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp := Resp{
		Language:    language,
		Translation: translationWord,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
