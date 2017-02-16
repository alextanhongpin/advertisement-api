package campaignsvc

// Add database logic here
import (
	"errors"

	"github.com/alextanhongpin/adsvc/common"
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	Create(Campaign) (string, error)
}

var (
	ErrInvalidId  = errors.New("Invalid Id")
	ErrorNotFound = errors.New("Not found")
)

type service struct{}

func (s service) All() ([]Campaign, error) {
	var campaigns []Campaign
	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("campaigns")
	iter := c.Find(nil).Iter()
	result := Campaign{}
	for iter.Next(&result) {
		campaigns = append(campaigns, result)
	}

	return campaigns, nil
}

func (s service) One(id string) (Campaign, error) {
	var campaign Campaign

	ds := common.NewDataStore()
	defer ds.Close()

	if !bson.IsObjectIdHex(id) {
		return campaign, ErrInvalidId
	}

	c := ds.C("campaigns")

	oid := bson.ObjectIdHex(id)

	err := c.FindId(oid).One(&campaign)

	if err != nil {
		return campaign, ErrorNotFound
	}
	return campaign, nil
}

func (s service) Create(cm Campaign) (string, error) {

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

func (s service) Delete(request interface{}) (bool, error) {
	// Verify id is ObjectId, otherwise bail
	req := request.(deleteRequest)
	if !bson.IsObjectIdHex(req.Id) {
		return false, ErrInvalidId
	}
	// Grab id
	oid := bson.ObjectIdHex(req.Id)
	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("campaigns")

	err := c.RemoveId(oid)

	if err != nil {
		return false, err
	}
	return true, nil
}
