package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"testing"

	"prot/entry-prot"
)

var filePath *string = flag.String("filePath", "", "FilePath")
var fileCategory *string = flag.String("fileCategory", "material", "[landscape / animal / material]")

const (
	UploadURL = "http://127.0.0.1:8687/file/upload"
)

func GetSourceParams() prot.FileUploadSource {
	pms := prot.FileUploadSource{
		UserID:   "1234567890",
		Category: *fileCategory,
	}

	_, fn := path.Split(*filePath)
	pms.FileName = fn

	return pms
}

func TestFileUpload(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	jsonBuf, err := json.Marshal(GetSourceParams())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	err = bodyWriter.WriteField("data", string(jsonBuf))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fileWriter, err := bodyWriter.CreateFormFile("file", *filePath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fh, err := os.Open(*filePath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(UploadURL, contentType, bodyBuf)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Error response StatusCode[%d], Content: %s", resp.StatusCode, resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(respBody))
}
