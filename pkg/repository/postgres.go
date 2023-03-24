package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type ServerConfig struct {
	Port string `json:"port"`
}

type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	SSLMode  string `json:"SSLMode"`
}

type Config struct {
	ServerConfig `json:"server_config"`
	DbConfig     `json:"db_config"`
}

func NewPostgresDb(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbConfig.Host,
		cfg.DbConfig.Port,
		cfg.DbConfig.Username,
		cfg.DbConfig.Password,
		cfg.DbConfig.DbName,
		cfg.DbConfig.SSLMode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
