package main

import (
	"fmt"
	"github.com/alextanhongpin/adsvc/adsvc"
	"github.com/alextanhongpin/adsvc/campaignsvc"
	"github.com/alextanhongpin/adsvc/common"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	config := common.GetConfig()
	router := httprouter.New()

	// Setup campaign service router
	router = campaignsvc.SetupRouter(router)
	router = adsvc.SetupRouter(router)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	// Serve static files
	// fs := http.FileServer(http.Dir("/static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("listening to port *%s.\npress ctrl + c to cancel", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, router))
}
