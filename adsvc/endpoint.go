package adsvc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alextanhongpin/adsvc/common"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

func (e Endpoint) All() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var advertisements []Advertisement

		ds := common.NewDataStore()
		defer ds.Close()

		c := ds.C("advertisements")

		iter := c.Find(nil).Iter()
		result := Advertisement{}
		for iter.Next(&result) {
			advertisements = append(advertisements, result)
		}

		j, err := json.Marshal(advertisements)

		if err != nil {
			panic(err)
		}

		ResponseWithJSON(w, j, 200)

	}
}

func (e Endpoint) Create() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ad := Advertisement{}
		json.NewDecoder(r.Body).Decode(&ad)

		ad.Id = bson.NewObjectId()

		ds := common.NewDataStore()
		defer ds.Close()

		c := ds.C("advertisements")

		c.Insert(ad)

		j, _ := json.Marshal(ad)

		ResponseWithJSON(w, j, 201)
	}
}

func (e Endpoint) One() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id := p.ByName("id")
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(404)
			return
		}
		oid := bson.ObjectIdHex(id)
		ad := Advertisement{}

		ds := common.NewDataStore()
		defer ds.Close()

		c := ds.C("advertisements")
		if err := c.FindId(oid).One(&ad); err != nil {
			w.WriteHeader(404)
			return
		}

		// Marshal provided interface into JSON structure
		j, _ := json.Marshal(ad)

		ResponseWithJSON(w, j, 201)
	}
}

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
