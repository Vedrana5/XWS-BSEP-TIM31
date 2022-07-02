package main

import (
	"post/module/startup"
	cf "post/module/startup/config"
)

func main() {
	config := cf.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
