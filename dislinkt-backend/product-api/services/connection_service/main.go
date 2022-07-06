package main

import (
	"connection/module/startup"
	cf "connection/module/startup/config"
)

const (
	DATABASE   = "posts"
	COLLECTION = "postsData"
)

func main() {
	config := cf.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
