package main

import (
	"fmt"
	"go-user-auth-service/internal/config"
	"go-user-auth-service/internal/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	addr := fmt.Sprintf(":%d", config.App.Port)
	r.Run(addr)
}
