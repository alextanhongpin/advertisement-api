package adsvc

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Advertisement struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name       string        `json:"name"`
	CreatedAt  time.Time     `json:"created_at"`
	ModifiedAt time.Time     `json:"modified_at"`
	// Title        string        `json:"title"`
	// StartAt      time.Time     `json:"start_at"`
	// EndAt        time.Time     `json:"end_at"`
	// Regions      []string      `json:"regions"`
	// ViewMaxCount int           `json:"view_max_count"`
	// ViewCount    int           `json:"view_count"`
	// Owner        string        `json:"owner"`
	// OwnerId      string        `json:"owner_id"`
	// Active       bool          `json:"active"`
	// Images       []string      `json:"images"`
	// CallToAction string        `json:"call_to_action"`
	// Links        []string      `json:"links`
}

type AdvertisementResource struct {
	Data Advertisement `json:"data"`
}
type AdvertisementCollection struct {
	Data []Advertisement `json:"data"`
}

type advertisementsRequest struct {
	Query string `json:"query,omitempty"`
}

type deleteRequest struct {
}
type deleteResponse struct {
	Ok    bool `json:"ok"`
	Error bool `json:"error,omitempty"`
}

type advertisementsResponse struct {
	Data []Advertisement `json:"data"`
	Err  string          `json:"err,omitempty"`
}
