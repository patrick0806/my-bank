package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type config struct {
	API_PORT string
}

func LoadEnvVars(env string) *config {
	if env != "PRODUCTION" {
		readEnvFile()
	}

	return &config{
		API_PORT: os.Getenv("API_PORT"),
	}
}

func readEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Falha ao carregar as variáveis de ambiente no modo dev: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro ao ler o arquivo .env: %v", err)
	}
}
