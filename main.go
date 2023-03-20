package main

import (
	"fmt"

	"github.com/danielbrazrocha/goportunities/config"
	router "github.com/danielbrazrocha/goportunities/router"
)

func main() {

	//Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		//panic(err)
		return
	}

	//Initialize router
	router.Initialize()
}
