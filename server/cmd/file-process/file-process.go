package main

import (
	"fmt"

	"conf"
	"entry"
	"router"
	_ "store"
)

func main() {
	router := router.NewHTTPRouter(conf.HTTPAddr())

	router.Register(entry.NewFileUpload())
	router.Register(entry.NewFileDownLoad())

	err := router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
