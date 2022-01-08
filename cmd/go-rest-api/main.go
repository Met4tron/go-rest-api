package main

import (
	"github.com/Met4tron/go-rest-api/pkg/config"
	"github.com/Met4tron/go-rest-api/pkg/infra/database"
	"github.com/Met4tron/go-rest-api/pkg/infra/http"
	"github.com/Met4tron/go-rest-api/pkg/logger"
)

func main() {
	config.LoadVariables()

	env := config.GetConfig()

	err := database.InitDB()

	if err != nil {
		logger.Fatal(err)
	}

	http.Init().Run(env.Server.Port)
}
