package config

import (
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbPassEsqSeq = "{password}"
	password     = "note-service-password"
)

type DB struct {
	DSN                string `json:"dsn"`
	MaxOpenConnections int32  `json:"max_open_connections"`
}

type GRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	DB   DB   `json:"db"`
	GRPC GRPC `json:"grpc"`
	HTTP HTTP `json:"http"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {

	dbDsn := strings.ReplaceAll(c.DB.DSN, dbPassEsqSeq, password)

	poolConfig, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.DB.MaxOpenConnections

	return poolConfig, nil
}

func (g *GRPC) GetAddress() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (h *HTTP) GetAddress() string {
	return net.JoinHostPort(h.Host, h.Port)
}
