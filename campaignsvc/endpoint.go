package campaignsvc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alextanhongpin/adsvc/common"
	"github.com/alextanhongpin/adsvc/helper"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

func (e Endpoint) All(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Index route for campaign service %s", svc.All())
		v, err := svc.All()
		if err != nil {
			panic(err)
		}
		res := CampaignCollection{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) GetOne(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		ps := ctx.Value("params").(httprouter.Params)

		v, err := svc.One(ps.ByName("id"))

		if err != nil {
			panic(err)
		}

		// Convert the objectId to string
		ot := oneTemplate{
			Id:          v.Id.Hex(),
			Name:        v.Name,
			CreatedAt:   v.CreatedAt,
			Description: v.Description,
		}

		common.RenderTemplate(w, "view-campaign", "base", ot)
	}
}

func (e Endpoint) One(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		p := ctx.Value("params").(httprouter.Params)

		fmt.Println(p)
		// ctx2.Value("hello")

		v, err := svc.One(p.ByName("id"))

		if err != nil {
			panic(err)
		}
		res := CampaignResource{
			Data: v,
		}

		j, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}
		helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) CreateGet(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Campaign{}
		common.RenderTemplate(w, "create-campaign", "base", c)
	}
}

func (e Endpoint) CreatePost(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// Define request
		req := Campaign{
			Name:        r.PostFormValue("campaign_name"),
			Description: r.PostFormValue("campaign_description"),
			StartAt:     time.Now(), //r.PostFormValue("start_at"),
			EndAt:       time.Now(), //r.PostFormValue("end_at"),
		}
		id, err := svc.Create(req)
		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/campaigns/"+id, 301)
	}
}

func (e Endpoint) Delete(svc service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ps := ctx.Value("params").(httprouter.Params)
		fmt.Println(ps)
		req := deleteRequest{
			Id: ps.ByName("id"),
		}
		v, err := svc.Delete(req)

		if err != nil {
			panic(err)
		}

		res := deleteResponse{
			Ok: v,
		}
		j, err := json.Marshal(res)

		if err != nil {
			panic(err)
		}

		// TODO: Update the status code
		helper.ResponseWithJSON(w, j, 200)
	}
}
