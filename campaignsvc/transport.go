package campaignsvc

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func SetupRouter(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	// Is it a good practice to initialize service this way?
	service := Service{}
	router.GET(wrap("/api/v1/campaigns", endpoint.All(service)))
	router.GET(wrap("/api/v1/campaigns/:id", endpoint.One(service)))
	router.DELETE(wrap("/api/v1/campaigns/:id", endpoint.Delete(service)))

	router.GET(wrap("/create/campaigns", endpoint.CreateGet(service)))
	router.POST(wrap("/create/campaigns", endpoint.CreatePost(service)))
	router.GET(wrap("/campaigns/:id", endpoint.GetOne(service)))
	return router
}

// func loggerMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Print("At logger middleware")
// 		next.ServeHTTP(w, r)
// 	})
// }

func wrap(p string, h func(http.ResponseWriter, *http.Request)) (string, httprouter.Handle) {
	return p, wrapHandler(alice.New().ThenFunc(h))
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
