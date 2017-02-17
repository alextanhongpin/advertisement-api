package campaignsvc

// Add database logic here
import (
	"errors"

	"github.com/alextanhongpin/adsvc/common"
	"github.com/alextanhongpin/adsvc/helper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorNotFound = errors.New("Not found")
)

type Service struct{}

// All returns a collection of campaigns
func (s Service) All(request interface{}) ([]Campaign, error) {
	// Define vars
	var iter *mgo.Iter

	// Request type assertion
	req := request.(allRequest)

	// Response resource
	res := []Campaign{}

	// Initialize collection
	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	// Handling empty string and nil is harder than I thought
	if req.Query == "" {
		iter = c.Find(nil).Iter()
	} else {
		iter = c.Find(req.Query).Iter()
	}

	var result Campaign
	for iter.Next(&result) {
		res = append(res, result)
	}
	return res, nil
}

// One returns a single campaign by id
func (s Service) One(request interface{}) (Campaign, error) {
	// Request type assertion
	req := request.(oneRequest)
	res := Campaign{}

	oid, err := helper.ValidateId(req.Id)
	if err != nil {
		return res, err
	}

	// Initialize collection
	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	if err = c.FindId(oid).One(&res); err != nil {
		return res, ErrorNotFound
	}
	return res, nil
}

func (s Service) Create(request interface{}) (string, error) {

	req := request.(createRequest)
	res := ""

	req.Data.Id = bson.NewObjectId()
	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("campaigns")

	if err := c.Insert(req.Data); err != nil {
		return res, err
	}
	res = req.Data.Id.Hex()
	return res, nil
}

func (s Service) Delete(request interface{}) (bool, error) {
	req := request.(deleteRequest)
	res := false
	// Verify id is ObjectId, otherwise bail
	oid, err := helper.ValidateId(req.Id)
	if err != nil {
		return res, err
	}

	// Initialize collection
	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	// Delete
	if err = c.RemoveId(oid); err != nil {
		return res, err
	}

	res = true

	return res, nil
}
