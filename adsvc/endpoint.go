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

		req := allRequest{
			Query: "",
		}

		v, err := svc.All(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := allResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) One(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps := helper.FetchParams(r)
		req := oneRequest{
			Id: ps.ByName("id"),
		}
		v, err := svc.One(req)

		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := oneResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		helper.ResponseWithJSON(w, j, 200)
	}
}

// Create with POST
func (e Endpoint) Create(svc Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Define a new model
		req := createRequest{
			Data: Advertisement{},
		}
		json.NewDecoder(r.Body).Decode(&req.Data)

		v, err := svc.Create(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		res := createResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
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
		ps := helper.FetchParams(r)
		req := deleteRequest{
			Id: ps.ByName("id"),
		}

		ok, err := svc.Delete(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		res := deleteResponse{
			Ok: ok,
		}
		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		helper.ResponseWithJSON(w, j, 201)
	}
}
