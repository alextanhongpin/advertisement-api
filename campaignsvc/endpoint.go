package campaignsvc

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Endpoint struct{}

func (e Endpoint) All(s Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "Index route for campaign service %s", s.All())
	}
}
