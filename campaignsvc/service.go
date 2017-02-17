package campaignsvc

// Add database logic here
import (
	"errors"

	"github.com/alextanhongpin/adsvc/common"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorInvalidId = errors.New("Invalid ID")
	ErrorNotFound  = errors.New("Not found")
)

type Service struct{}

func (s Service) All(request interface{}) ([]Campaign, error) {
	// Handle request
	req := request.(allRequest)

	var campaigns []Campaign

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	iter := c.Find(req.Query).Iter()
	res := Campaign{}

	for iter.Next(&result) {
		campaigns = append(campaigns, res)
	}
	return campaigns, nil
}

func (s Service) One(request interface{}) (Campaign, error) {
	req := request.(oneRequest)

	var campaign Campaign
	if !bson.IsObjectIdHex(req.Id) {
		return campaign, ErrorInvalidId
	}
	oid := bson.ObjectIdHex(req.Id)

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	err := c.FindId(oid).One(&campaign)

	if err != nil {
		return campaign, ErrorNotFound
	}
	return campaign, nil
}

func (s Service) Create(request interface{}) (string, error) {
	req := request.(createRequest)
	cm.Id = bson.NewObjectId()
	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("campaigns")

	err := c.Insert(cm)

	if err != nil {
		return "", err
	}
	return cm.Id.Hex(), nil
}

func (s Service) Delete(request interface{}) (bool, error) {
	req := request.(deleteRequest)
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(req.Id) {
		return false, ErrorInvalidId
	}
	// Grab id
	oid := bson.ObjectIdHex(req.Id)

	// Initialize collection
	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("campaigns")

	// Action
	err := c.RemoveId(oid)

	if err != nil {
		return false, err
	}
	return true, nil
}
