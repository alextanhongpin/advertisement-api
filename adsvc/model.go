package adsvc

import (
	// "github.com/asaskevich/govalidator"
	// "github.com/google/jsonapi"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Advertisement struct {
	Id         bson.ObjectId `valid:"-" bson:"_id,omitempty" json:"id"`
	Name       string        `valid:"-" json:"name"`
	CreatedAt  time.Time     `valid:"-" json:"created_at"`
	ModifiedAt time.Time     `valid:"-" json:"modified_at"`
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
