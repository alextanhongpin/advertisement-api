package adsvc

import (
	"github.com/alextanhongpin/adsvc/common"
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	All() []Advertisement
}

// ErrEmpty is returned when an input string is empty.
var (
	ErrInvalidId  = errors.New("Invalid Id")
	ErrorNotFound = errors.New("Not found")
)

type service struct{}

// All returns a list of advertisements
func (s service) All(request interface{}) ([]Advertisement, error) {

	var advertisements []Advertisement

	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("advertisements")
	iter := c.Find(nil).Iter()
	result := Advertisement{}
	for iter.Next(&result) {
		advertisements = append(advertisements, result)
	}

	return advertisements, nil
}

func (s service) One(id string) (Advertisement, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)
	ad := Advertisement{}

	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("advertisements")
	err := c.FindId(oid).One(&ad)

	if err != nil {
		return nil, ErrorNotFound
	}
	return ad, nil
}

// Create accepts a new advertisement model
// and returns the created resource and an error
func (s service) Create(ad Advertisement) (Advertisement, error) {

		ds := common.NewDataStore()
		defer ds.Close()

		c := ds.C("advertisements")

		err := c.Insert(ad)

		if err != nil {
			return nil, err
		}
		return ad, nil
	}
}
