package main

import (
	"flag"

	"github.com/patrick0806/my-bank/config"
)

func main() {
	env := flag.String("env", "default", "Environment to run api")
	flag.Parse()
	configs := config.LoadEnvVars(*env)
	println(configs.API_PORT)
}
