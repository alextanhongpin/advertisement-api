package helper

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorInvalidId = errors.New("Invalid ID")
)

// ValidateId returns a valid bsonObject id or an error if
// the string is not a valid mongoid
func ValidateId(id string) (bson.ObjectId, error) {
	// Cannot return nil
	oid := bson.NewObjectId()
	// String cannot be empty
	if id == "" {
		return oid, ErrorInvalidId
	}
	// Not a valid object id
	if !bson.IsObjectIdHex(id) {
		return oid, ErrorInvalidId
	}
	// Valid, parse id
	oid = bson.ObjectIdHex(id)
	return oid, nil
}
