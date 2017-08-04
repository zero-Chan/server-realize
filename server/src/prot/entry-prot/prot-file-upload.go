package prot

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
)

type FileUploadSource struct {
	UserID   string `json:"UserID"`
	Category string `json:"Category"`
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

	case src.Category != ImageCategory_Animal && src.Category != ImageCategory_Material && src.Category != ImageCategory_Landscape:
		return fmt.Errorf("invalid params: [Category] is [%s], data:%s", src.Category, data)

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
