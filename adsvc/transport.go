package adsvc

import (
	"github.com/julienschmidt/httprouter"
)

func SetupRouter(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	router.GET("/advertisements", endpoint.All())
	router.GET("/advertisements/:id", endpoint.One())
	router.POST("/advertisements", endpoint.Create())
	return router
}
