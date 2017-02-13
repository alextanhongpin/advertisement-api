package campaignsvc

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Campaign struct {
	Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name         string        `json:"name"` // The name of the campaign as seen by admin
	CreatedAt    time.Time     `json:"created_at"`
	ModifiedAt   time.Time     `json:"modified_at"`
	Title        string        `json:"title"`          // The title of the campaign as seen by end user
	StartAt      time.Time     `json:"start_at"`       // The start date of the campaign
	EndAt        time.Time     `json:"end_at"`         // The end date of the campaign
	Regions      []string      `json:"regions"`        // The participating regions
	ViewMaxCount int           `json:"view_max_count"` // The maximum view count before the ads expired
	ViewCount    int           `json:"view_count"`     // The current view count
	Owner        string        `json:"owner"`          // The campaign owner
	OwnerId      string        `json:"owner_id"`
	Active       bool          `json:"active"` // The state of the campaign
}
