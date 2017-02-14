package common

import (
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	config := GetConfig()
	s, err := mgo.Dial(config.MongoURI)

	if err != nil {
		panic(err)
	}
	return s
}
