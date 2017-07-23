package entry

import (
	//	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	//	"strings"
)

type FileDownload struct {
	method string
	path   string
	src    FileDownloadSource
	res    FileDownloadResult
}

func NewFileDownLoad() *FileDownload {
	entry := &FileDownload{
		method: http.MethodGet,
		path:   "/file/down",
		src:    CreateFileDownloadSource(),
		res:    CreateFileDownLoadResult(),
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

type FileDownloadSource struct {
	FilePath string
}

func CreateFileDownloadSource() FileDownloadSource {
	src := FileDownloadSource{}
	return src
}

func (src *FileDownloadSource) Get(req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return fmt.Errorf("Parse URL Form error: %s", err)
	}

	src.FilePath = req.Form.Get("file")

	switch {
	case len(src.FilePath) == 0:
		return fmt.Errorf("Invalid [FilePath] empty. url-form: %+v", req.Form)
	}

	return nil
}

type FileDownloadResult struct {
	File multipart.File
}

func CreateFileDownLoadResult() FileDownloadResult {
	res := FileDownloadResult{}
	return res
}

func (res *FileDownloadResult) Response(w http.ResponseWriter) error {

	w.WriteHeader(http.StatusOK)

	return nil
}
