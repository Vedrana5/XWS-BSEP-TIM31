package main

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/startup"
	cf "github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/startup/config"
)

func main() {
	config := cf.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
