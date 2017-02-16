package adsvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	"github.com/alextanhongpin/adsvc/common"
	"github.com/alextanhongpin/adsvc/helper"
)

type Endpoint struct{}

func (e Endpoint) Index(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ad := Advertisement{}
		common.RenderTemplate(w, "index", "base", ad)
	}
}

func (e Endpoint) All(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := advertisementsRequest{}

		v, err := svc.All(req)
		if err != nil {
			panic(err)
		}

		res := advertisementsResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		helper.ResponseWithJSON(w, j, 200)
	}
}

// Create with POST
func (e Endpoint) Create(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Define a new model
		ad := Advertisement{}
		json.NewDecoder(r.Body).Decode(&ad)
		ad.Id = bson.NewObjectId()

		v, err := svc.Create(ad)
		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		helper.ResponseWithJSON(w, j, 201)
	}
}

func (e Endpoint) Success() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "success")
	}
}

// Create with Form
func (e Endpoint) CreateForm(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		r.ParseForm()

		// fmt.Fprintf(w, "{message: %q}", r.PostFormValue("name"))
		http.Redirect(w, r, "/success", 301)
	}
}

func (e Endpoint) Delete(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := ps.ByName("id")

		ok, err := svc.Delete(id)

		if err != nil {
			panic(err)
		}
		res := deleteResponse{
			Ok: ok,
		}
		j, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		helper.ResponseWithJSON(w, j, 201)
	}
}
