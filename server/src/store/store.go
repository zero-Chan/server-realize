package store

import (
	"fmt"
	"os"

	"conf"
	"prot/project-prot"
	"store/base"
)

var (
	defaultStore = CreateStore()
)

func GlobalStore() *Store {
	return &defaultStore
}

type Store struct {
	MongoServer *base.MongoStore
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
		MongoServer: base.NewMongoStore(),
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

	err := store.MongoServer.Dial(conf.String())
	if err != nil {
		return fmt.Errorf("Store.mongoUp fail: %s", err)
	}

	return nil
}

func (store *Store) GetControllor(info prot.ProjectInfoProt) StoreController {
	ctl := CreateStoreController(StoreControllerConfig{
		// projectInfo
		proInfo: &info,

		// Conns
		MongoServerConn: store.MongoServer.Session(),
	})
	return ctl
}
