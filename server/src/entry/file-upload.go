package entry

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"conf"
	"prot/entry-prot"
	proProt "prot/project-prot"
	"store"
)

type FileUpload struct {
	method      string
	path        string
	contentType string
	src         prot.FileUploadSource
	res         prot.FileUploadResult
}

func NewFileUpload() *FileUpload {
	entry := &FileUpload{
		method:      http.MethodPost,
		path:        "/file/upload",
		contentType: "multipart/form-data",
		src:         prot.CreateFileUploadSource(),
		res:         prot.CreateFileUploadResult(),
	}

	return entry
}

func (this *FileUpload) Path() string {
	return this.path
}

func (this *FileUpload) RouterRule(req *http.Request) bool {
	if req.Method != this.method {
		return false
	}

	if req.URL.Path != this.path {
		return false
	}

	cl := strings.Split(req.Header.Get("Content-Type"), ";")
	if cl[0] != this.contentType {
		return false
	}

	return true
}

func (this *FileUpload) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := this.src.Get(req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileID := bson.NewObjectId().String()
	filep, err := os.OpenFile(conf.FilesPath()+"/"+fileID, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(filep, this.src.File)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filep.Close()

	// 防止资源竞争，不让两个程序操作同一个文件
	targetFile := conf.FilesPath() + "/" + this.src.FileName
	err = os.Rename(conf.FilesPath()+"/"+fileID, targetFile)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	this.src.FileName = targetFile
	this.res.FilePath = targetFile

	// save file info to db
	storeCtl := store.GlobalStore().GetControllor(proProt.ProjectInfoProt{})

	err = storeCtl.EntityFileStore().SaveFileInfo(&this.src)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = this.res.Response(w)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
