package config

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"os"
	"strings"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func LoadConfig() (*core.Config, error) {
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
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "json")

	certPool, err := loadCACertPool(viper.GetString("ca_cert_path"))
	if err != nil {
		return nil, err
	}

	vaultToken, err := ioutil.ReadFile(viper.GetString("vault_token_path"))
	if err != nil {
		return nil, err
	}

	certify, err := loadCertify(viper.GetString("vault_addr"), viper.GetString("vault_pki"), viper.GetString("vault_role"), viper.GetString("vault_cn"), strings.TrimSpace(string(vaultToken)))

	tlsConfig := &tls.Config{
		GetClientCertificate: certify.GetClientCertificate,
		RootCAs:              certPool,
	}

	grpcConn, err := grpc.Dial(viper.GetString("grpc_host_auth"), grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))

	logLevel := core.LogLevelInfo
	if viper.GetString("log_level") == "debug" {
		logLevel = core.LogLevelDebug
	}

	cfg := &core.Config{
		// http server
		HttpPort: viper.GetString("http_port"),
		// logging
		LogLevel:  logLevel,
		LogFormat: core.LogFormat(viper.GetString("log_format")),
		// grpc
		GrpcConn: grpcConn,
		// postgres
		PgHost:       viper.GetString("pg_host"),
		PgPort:       viper.GetString("pg_port"),
		PgUsername:   viper.GetString("pg_username"),
		PgPassword:   viper.GetString("pg_password"),
		PgDatabase:   viper.GetString("pg_database"),
		PgParameters: viper.GetString("pg_parameters"),
	}

	return cfg, nil
}
