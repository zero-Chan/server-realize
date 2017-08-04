package entry

import (
	"fmt"
	"net/http"
	"prot/entry-prot"
	"strings"
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

	// TODO
	fmt.Printf("Source: %+v\n", this.src)

	err = this.res.Response(w)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
