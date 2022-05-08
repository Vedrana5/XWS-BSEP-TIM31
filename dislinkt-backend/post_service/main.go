package main

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/post_service/startup"
	cfg "github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/post_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
