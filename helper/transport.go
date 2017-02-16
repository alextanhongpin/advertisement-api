package helper

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/vnd.api+json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
func FetchParams(r *http.Request) httprouter.Params {
	ctx := r.Context()
	return ctx.Value("params").(httprouter.Params)
}
