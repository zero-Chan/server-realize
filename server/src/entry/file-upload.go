package entry

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

type FileUpload struct {
	method      string
	path        string
	contentType string
	src         FileUploadSource
	res         FileUploadResult
}

func NewFileUpload() *FileUpload {
	entry := &FileUpload{
		method:      http.MethodPost,
		path:        "/file/upload",
		contentType: "multipart/form-data",
		src:         CreateFileUploadSource(),
		res:         CreateFileUploadResult(),
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

type FileUploadSource struct {
	UserID   string `json:"UserID"`
	FileName string `json:"FileName"`

	File multipart.File
}

func CreateFileUploadSource() FileUploadSource {
	src := FileUploadSource{}

	return src
}

func (src *FileUploadSource) Get(req *http.Request) error {
	err := req.ParseMultipartForm(http.DefaultMaxHeaderBytes)
	if err != nil {
		return err
	}

	data := req.FormValue("data")

	err = json.Unmarshal([]byte(data), src)
	if err != nil {
		return err
	}

	switch {
	case len(src.UserID) == 0:
		return fmt.Errorf("invalid params: [UserID] is empty. data:%s", data)

	case len(src.FileName) == 0:
		return fmt.Errorf("invalid params: [FileName] is empty. data:%s", data)
	}

	src.File, _, err = req.FormFile("file")
	if err != nil {
		return fmt.Errorf("read file error: %s", err)
	}

	return nil
}

type FileUploadResult struct {
	FilePath string `json:"FilePath"`
}

func CreateFileUploadResult() FileUploadResult {
	res := FileUploadResult{}

	return res
}

func (res *FileUploadResult) Response(w http.ResponseWriter) error {
	body := make(map[string]interface{})
	body["message"] = res

	b, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("Json Marshal fail: %s", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(b)
	return err
}
