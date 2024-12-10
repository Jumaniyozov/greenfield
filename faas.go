package faas

import (
	"github.com/jumaniyozov/greenfield/handlers/rest"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
