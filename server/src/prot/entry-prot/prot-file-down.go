package prot

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

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
