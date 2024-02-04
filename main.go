package main

import (
	"flag"

	"github.com/patrick0806/my-bank/api"
	"github.com/patrick0806/my-bank/config"
	"github.com/patrick0806/my-bank/config/database"
)

func main() {
	env := flag.String("env", "PRODUCTION", "Environment to run api")
	flag.Parse()
	configs := config.LoadEnvVars(*env)
	database.GetDB()
	api.Run(configs.API_PORT)
}
