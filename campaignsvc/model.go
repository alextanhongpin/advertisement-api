package campaignsvc

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	Campaign struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name         string        `json:"name"` // The name of the campaign as seen by admin
		CreatedAt    time.Time     `json:"created_at"`
		UpdatedAt    time.Time     `json:"updated_at"`
		StartAt      time.Time     `json:"start_at"` // The start date of the campaign
		EndAt        time.Time     `json:"end_at"`   // The end date of the campaign
		Description  string        `json:"description"`
		Title        string        `json:"title"`          // The title of the campaign as seen by end user
		Region       string        `json:"region"`         // The participating regions
		ViewMaxCount int           `json:"view_max_count"` // The maximum view count before the ads expired
		ViewCount    int           `json:"view_count"`     // The current view count
		Owner        string        `json:"owner"`          // The campaign owner
		OwnerId      string        `json:"owner_id"`
		Active       bool          `json:"active"` // The state of the campaign
		Country      string        `json:"country"`
		City         string        `json:"city"`
	}
	// CampaignCollection struct {
	// 	Data []Campaign `json:"data"`
	// }
	// CampaignResource struct {
	// 	Data Campaign `json:"data"`
	// }

	allRequest struct {
		Query string `json:"query"`
	}
	allResponse struct {
		Data []Campaign `json:"data"`
	}

	oneRequest struct {
		Id string `json:"id"`
	}
	oneResponse struct {
		Data Campaign `json:"data"`
	}

	createRequest struct {
		Data Campaign `json:"data"`
	}

	createResponse struct {
		Id string `json:"id"`
	}

	updateRequest  struct{}
	updateResponse struct{}

	deleteRequest struct {
		Id string `json:"id"`
	}
	deleteResponse struct {
		Ok bool `json:"ok"`
	}

	oneTemplate struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
	}
)

// groupby Grouping parameter - link, recipient, domain, country, region, city, month, day, hour, minute, daily_hour.
