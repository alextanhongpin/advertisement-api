package adsvc

import (
	"github.com/alextanhongpin/adsvc/common"
)

type Service interface {
	All() []Advertisement
}

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
