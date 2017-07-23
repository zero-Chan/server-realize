package conf

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	defaultConfig         = CreateConfig()
	confFile      *string = flag.String("filename", "conf/conf.json", "Load Configure File *.json")
)

type Config struct {
	HTTPAddr    *string          `json:"HttpAddr"`
	MongoServer *MongoServerConf `json:"MongoServer"`
}

func HTTPAddr() string {
	if defaultConfig.HTTPAddr == nil {
		return ""
	}

	return *defaultConfig.HTTPAddr
}

func MongoServer() (MongoServerConf, bool) {
	if defaultConfig.MongoServer == nil {
		return CreateMongoServerConf(), false
	}

	return *defaultConfig.MongoServer, true
}

func (conf *Config) LoadByJSON(data []byte) error {
	err := json.Unmarshal(data, conf)
	if err != nil {
		return err
	}

	switch {
	case conf.HTTPAddr == nil:
		return fmt.Errorf("Conf.HTTPAddr must to init.")
	case conf.MongoServer != nil:
		if err := conf.MongoServer.Vaild(); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	flag.Parse()

	buf, err := LoadFile()
	if err != nil {
		fmt.Printf("Load Conf File fail: %s\n", err)
		os.Exit(1)
	}

	err = defaultConfig.LoadByJSON(buf)
	if err != nil {
		fmt.Printf("Parse Config(file=%s) By jSON fail: %s\n", *confFile, err)
		os.Exit(1)
	}
}

func CreateConfig() Config {
	conf := Config{}
	return conf
}

func LoadFile(file ...string) ([]byte, error) {
	var f string

	if len(file) > 1 {
		return nil, fmt.Errorf("Only Can Load a config file.")
	} else if len(file) == 0 {
		f = *confFile
	} else if len(file) == 1 {
		f = file[0]
	}

	filep, err := os.OpenFile(f, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("Open File[%s] fail: %s", f, err)
	}

	buf := bytes.NewBuffer(nil)

	// the copy func use 'append' way to buf.
	_, err = io.Copy(buf, filep)
	if err != nil {
		return nil, fmt.Errorf("Copy File[%s] data fail: %s", f, err)
	}

	return buf.Bytes(), nil
}

type MongoServerConf struct {
	Host       string `json:"Host"`
	Port       int64  `json:"Port"`
	DB         string `json:"Db"`
	Collection string `json:"Collection"`
}

func CreateMongoServerConf() MongoServerConf {
	conf := MongoServerConf{}
	return conf
}

func (conf *MongoServerConf) Vaild() error {
	switch {
	case len(conf.Host) == 0:
		return fmt.Errorf("MongoServerConf.Host empty")
	case conf.Port <= 0:
		return fmt.Errorf("MongoServerConf.Port[%d] invalid.", conf.Port)
	}

	return nil
}

func (conf *MongoServerConf) String() string {
	url := fmt.Sprintf("mongodb://%s:%d", conf.Host, conf.Port)
	return url
}
