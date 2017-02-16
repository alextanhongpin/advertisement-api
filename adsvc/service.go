package adsvc

import (
	"errors"
	"github.com/alextanhongpin/adsvc/common"
	"gopkg.in/mgo.v2/bson"
)

// ErrEmpty is returned when an input string is empty.
var (
	ErrInvalidId  = errors.New("Invalid Id")
	ErrorNotFound = errors.New("Not found")
)

type Service struct{}

// All returns a list of advertisements
func (s Service) All(request interface{}) ([]Advertisement, error) {

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

func (s Service) One(id string) (Advertisement, error) {
	ad := Advertisement{}
	if !bson.IsObjectIdHex(id) {
		return ad, ErrInvalidId
	}
	oid := bson.ObjectIdHex(id)

	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("advertisements")
	err := c.FindId(oid).One(&ad)

	if err != nil {
		return ad, ErrorNotFound
	}
	return ad, nil
}

// Create accepts a new advertisement model
// and returns the created resource and an error
func (s Service) Create(ad Advertisement) (Advertisement, error) {

	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("advertisements")

	err := c.Insert(ad)

	if err != nil {
		return ad, err
	}
	return ad, nil
}

func (s Service) Delete(id string) (bool, error) {
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return false, ErrInvalidId
	}

	// Grab id
	oid := bson.ObjectIdHex(id)
	ds := common.NewDataStore()
	defer ds.Close()

	c := ds.C("advertisements")

	err := c.RemoveId(oid)

	if err != nil {
		return false, err
	}
	return true, nil
}
