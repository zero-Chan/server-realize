package entry

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"prot/entry-prot"
)

type FileStore interface {
	SaveFileInfo(*prot.FileUploadSource) error
}

type FileStoreConfig struct {
}

// Mgo DB
const (
	MgoDBName         = "ObjectStore"
	MgoCollectionName = "FileInfo"
)

type MgoFileStoreConfig struct {
	FileStoreConfig
	Conn *mgo.Session
}

type MgoFileStore struct {
	cfg *MgoFileStoreConfig
}

func CreateMgoFileStore(cfg MgoFileStoreConfig) MgoFileStore {
	store := MgoFileStore{
		cfg: &cfg,
	}
	return store
}

func NewMgoFileStore(cfg MgoFileStoreConfig) *MgoFileStore {
	s := CreateMgoFileStore(cfg)
	return &s
}

func (store *MgoFileStore) SaveFileInfo(src *prot.FileUploadSource) error {
	if src == nil {
		return fmt.Errorf("SaveFileInfo get nil params.")
	}

	coll := store.cfg.Conn.DB(MgoDBName).C(MgoCollectionName)

	upSelector := map[string]interface{}{
		"_id": src.FileName,
	}

	upInfo := map[string]interface{}{
		"_id":        src.FileName,
		"UserID":     src.UserID,
		"UpdateTime": bson.Now(),
		"Category":   src.Category,
	}

	_, err := coll.Upsert(upSelector, upInfo)
	if err != nil {
		return fmt.Errorf("SaveFileInfo to mongodb fail: %s", err)
	}

	return nil
}
