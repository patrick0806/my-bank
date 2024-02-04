package main

import (
	"flag"

	"github.com/patrick0806/my-bank/api"
	"github.com/patrick0806/my-bank/config"
	"github.com/patrick0806/my-bank/config/database"
	_ "github.com/patrick0806/my-bank/docs"
)

// @title My Bank API
// @version 1.0
// @description A api for my bank operations
// @BasePath /api/v1
// @contact.name My bank support
// @contact.url mybank.com.br
// @contact.email patrick@mybank.com.br

/*
// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
*/
func main() {
	env := flag.String("env", "PRODUCTION", "Environment to run api")
	flag.Parse()
	configs := config.LoadEnvVars(*env)
	database.GetDB()
	api.Run(configs.API_PORT)
}
