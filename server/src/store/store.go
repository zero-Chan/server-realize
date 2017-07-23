package store

import (
	"fmt"
	"os"

	"conf"
	"store/base"
)

var (
	defaultStore = CreateStore()
)

type Store struct {
	Mgo *base.MongoStore
}

func init() {
	err := defaultStore.Up()
	if err != nil {
		fmt.Printf("Store init fail: %s\n", err)
		os.Exit(1)
	}
}

func CreateStore() Store {
	store := Store{
		Mgo: base.NewMongoStore(),
	}

	return store
}

func (store *Store) Up() error {
	err := store.mgoUp()
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) mgoUp() error {
	conf, exist := conf.MongoServer()
	if !exist {
		return fmt.Errorf("Conf.MongoServer never exist.")
	}

	err := store.Mgo.Dial(conf.String())
	if err != nil {
		return fmt.Errorf("Store.mongoUp fail: %s", err)
	}

	return nil
}

func (store *Store) Copy() *Store {
	news := CreateStore()
	news.Mgo = store.Mgo.Copy()

	return &news
}
