package main

import (
	"post/module/startup"
	cf "post/module/startup/config"
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
