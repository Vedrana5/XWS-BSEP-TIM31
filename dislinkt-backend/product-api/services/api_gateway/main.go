package main

import (
	"fmt"
	startup "gateway/module/startup"
	config "gateway/module/startup/config"
)

func main() {
	config := config.NewConfig()
	fmt.Printf("PORT GATEWAYA JE" + config.Port)
	server := startup.NewServer(config)
	server.Start()
}
