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
	var iter *mgo.Iter
	req := request.(allRequest)
	res := []Advertisement{}

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("advertisements")

	if req.Query == "" {
		iter = c.Find(nil).Iter()
	} else {
		iter = c.Find(req.Query).Iter()
	}
	result := Advertisement{}
	for iter.Next(&result) {
		res = append(res, result)
	}

	return res, nil
}

func (s Service) One(request interface{}) (Advertisement, error) {
	req := request.(oneRequest)
	res := Advertisement{}

	oid, err := helper.ValidateId(req.Id)
	if err != nil {
		return res, err
	}

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("advertisements")

	if err := c.FindId(oid).One(&res); err != nil {
		return res, ErrorNotFound
	}

	return res, nil
}

// Create accepts a new advertisement model
// and returns the created resource and an error
func (s Service) Create(request interface{}) (bool, error) {
	req := request.(createRequest)
	res := false

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("advertisements")

	if err := c.Insert(req.Data); err != nil {
		return res, err
	}
	res = true
	return res, nil
}

func (s Service) Delete(request interface{}) (bool, error) {
	req := request.(deleteRequest)
	res := false

	oid, err := helper.ValidateId(req.Id)
	if err != nil {
		return res, err
	}

	ds := common.NewDataStore()
	defer ds.Close()
	c := ds.C("advertisements")

	if err := c.RemoveId(oid); err != nil {
		return res, err
	}

	res = true
	return res, nil
}
