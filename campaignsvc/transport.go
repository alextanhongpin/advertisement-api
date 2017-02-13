package campaignsvc

import (
	"github.com/julienschmidt/httprouter"
)

func SetupRouter(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	service := Service{}
	router.GET("/campaign", endpoint.All(service))
	return router
}
