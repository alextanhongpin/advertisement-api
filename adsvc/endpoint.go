package adsvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

type Endpoint struct{}

func (e Endpoint) All(svc service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

		ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) Create(svc service) httprouter.Handle {
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

		ResponseWithJSON(w, j, 201)
	}
}

func (e Endpoint) Delete(svc service) httprouter.Handle {
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
		ResponseWithJSON(w, j, 201)
	}
}

// func (e Endpoint) One() httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 		id := p.ByName("id")
// 		if !bson.IsObjectIdHex(id) {
// 			w.WriteHeader(404)
// 			return
// 		}
// 		oid := bson.ObjectIdHex(id)
// 		ad := Advertisement{}

// 		ds := common.NewDataStore()
// 		defer ds.Close()

// 		c := ds.C("advertisements")
// 		if err := c.FindId(oid).One(&ad); err != nil {
// 			w.WriteHeader(404)
// 			return
// 		}

// 		// Marshal provided interface into JSON structure
// 		j, _ := json.Marshal(ad)

// 		ResponseWithJSON(w, j, 201)
// 	}
// }

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
