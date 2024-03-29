package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/patrick0806/my-bank/api"
	"github.com/patrick0806/my-bank/config"
	_ "github.com/patrick0806/my-bank/docs"
	"github.com/patrick0806/my-bank/pkg/queues"
)

// @title My Bank API
// @version 1.0
// @description A api for my bank operations
// @BasePath /api/v1
// @contact.name My bank support
// @contact.url mybank.com.br
// @contact.email patrick@mybank.com.br
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name accessToken
/*
// @externalDocs.description  OpenAPI
*/
func main() {
	env := flag.String("env", "PRODUCTION", "Environment to run api")
	flag.Parse()
	configs := config.LoadEnvVars(*env)

	db, err := sql.Open(os.Getenv("DB_DRIVER"), fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	if err != nil {
		log.Fatalf("Failt to open connection with database", err)
	}
	defer db.Close()
	queue := queues.NewTransactionQueue(db)
	// Iniciar a goroutine para processar as transações em segundo plano
	go queue.ProcessTransactions()

	api.Run(configs.API_PORT, db, queue)
}
