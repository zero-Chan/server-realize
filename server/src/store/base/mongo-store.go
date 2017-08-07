package base

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	session *mgo.Session
}

func CreateMongoStore() MongoStore {
	store := MongoStore{}
	return store
}

func NewMongoStore() *MongoStore {
	store := CreateMongoStore()

	return &store
}

func (store *MongoStore) Dial(url string) error {
	var err error

	store.session, err = mgo.Dial(url)
	if err != nil {
		return fmt.Errorf("Mongo Dial fail: %s", err)
	}

	return store.session.Ping()
}

func (store *MongoStore) Session() *mgo.Session {
	return store.session
}
