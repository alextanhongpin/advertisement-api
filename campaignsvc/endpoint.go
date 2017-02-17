package campaignsvc

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alextanhongpin/adsvc/common"
	"github.com/alextanhongpin/adsvc/helper"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

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

func (e Endpoint) GetOne(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ps := ctx.Value("params").(httprouter.Params)

		req := oneRequest{
			Id: ps.ByName("id"),
		}

		v, err := svc.One(req)
		if err != nil {
			// This is a template, probably redirect to 404 page?
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		// Convert the objectId to string
		res := oneTemplate{
			Id:          v.Id.Hex(),
			Name:        v.Name,
			CreatedAt:   v.CreatedAt,
			Description: v.Description,
		}

		common.RenderTemplate(w, "view-campaign", "base", res)
	}
}

func (e Endpoint) One(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ps := ctx.Value("params").(httprouter.Params)

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

func (e Endpoint) CreateGet(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := Campaign{}
		common.RenderTemplate(w, "create-campaign", "base", c)
	}
}

func (e Endpoint) CreatePost(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		req := createRequest{
			Data: Campaign{
				Name:        r.PostFormValue("campaign_name"),
				Description: r.PostFormValue("campaign_description"),
				StartAt:     time.Now(), //r.PostFormValue("start_at"),
				EndAt:       time.Now(), //r.PostFormValue("end_at"),
			},
		}
		id, err := svc.Create(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		http.Redirect(w, r, "/campaigns/"+id, 301)
	}
}

func (e Endpoint) Delete(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ps := ctx.Value("params").(httprouter.Params)
		req := deleteRequest{
			Id: ps.ByName("id"),
		}

		v, err := svc.Delete(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := deleteResponse{
			Ok: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		// TODO: Update the status code
		helper.ResponseWithJSON(w, j, 200)
	}
}
