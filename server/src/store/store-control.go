package store

import (
	"gopkg.in/mgo.v2"

	"prot/project-prot"
	"store/entry"
)

type StoreControllerConfig struct {
	// projectInfo
	proInfo *prot.ProjectInfoProt

	// Conns
	MongoServerConn *mgo.Session
}

type StoreController struct {
	cfg *StoreControllerConfig

	entryFileStore entry.FileStore
}

func CreateStoreController(cfg StoreControllerConfig) StoreController {
	ctl := StoreController{
		cfg: &cfg,
	}
	return ctl
}

func NewStoreController(cfg StoreControllerConfig) *StoreController {
	ctl := CreateStoreController(cfg)
	return &ctl
}

func (ctl *StoreController) EntityFileStore() entry.FileStore {
	// Use MongoServerConn
	if ctl.entryFileStore == nil {
		ctl.entryFileStore = entry.NewMgoFileStore(entry.MgoFileStoreConfig{
			Conn:            ctl.cfg.MongoServerConn.Copy(),
			FileStoreConfig: entry.FileStoreConfig{},
		})
	}

	return ctl.entryFileStore
}
