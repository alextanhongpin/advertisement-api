package common

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

type DataStore struct {
	session *mgo.Session
}

func (d *DataStore) Close() {
	d.session.Close()
}

func init() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
}

func (d *DataStore) C(name string) *mgo.Collection {
	config := GetConfig()
	return d.session.DB(config.MongoDB).C(name)
}

func NewDataStore() *DataStore {
	ds := &DataStore{
		session: session.Copy(),
	}
	return ds
}
