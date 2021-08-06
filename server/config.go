package server

import (
	"log"
	"os"
	"strings"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func loadConfig() *core.Config {
	envPath := os.Getenv("ENV_PATH")
	if len(envPath) == 0 {
		envPath = ".env"
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("failed to load .env: %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("http_port", "8080")

	return &core.Config{
		HttpPort: viper.GetString("http_port"),
	}
}
