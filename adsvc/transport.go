package adsvc

import (
	"github.com/julienschmidt/httprouter"
)

func SetupRouter(router *httprouter.Router) *httprouter.Router {
	endpoint := Endpoint{}
	svc := service{}
	router.GET("/advertisements", endpoint.All(svc))
	// router.GET("/advertisements/:id", endpoint.One(svc))
	// router.POST("/advertisements", endpoint.Create(svc))
	return router
}
