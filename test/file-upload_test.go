package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

const (
	confFile = "./conf/conf-file-upload.json"
)

func LoadConfFile(file string) ([]byte, error) {
	fp, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("Open Conf file[%s] fail. error: %s", file, err)
	}

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, fp)
	if err != nil {
		return nil, fmt.Errorf("Get Conf file[%s] Content fail. error: %s", file, err)
	}

	return buf.Bytes(), nil
}

func TestFileUpload(t *testing.T) {
	buf, err := LoadConfFile(confFile)
	if err != nil {
		t.Errorf(err)
		t.FailNow()
	}

	err = json.Unmarshal(buf)
	if err != nil {

	}
}
