package entry

import (
	"fmt"
	"net/http"
	"prot/entry-prot"
)

type FileDownload struct {
	method string
	path   string
	src    prot.FileDownloadSource
	res    prot.FileDownloadResult
}

func NewFileDownLoad() *FileDownload {
	entry := &FileDownload{
		method: http.MethodGet,
		path:   "/file/down",
		src:    prot.CreateFileDownloadSource(),
		res:    prot.CreateFileDownLoadResult(),
	}

	return entry
}

func (this *FileDownload) Path() string {
	return this.path
}

func (this *FileDownload) RouterRule(req *http.Request) bool {
	if req.Method != this.method {
		return false
	}

	if req.URL.Path != this.path {
		return false
	}

	return true
}

func (this *FileDownload) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
