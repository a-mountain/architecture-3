package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConnectionConfig struct {
	DbName   string
	User     string
	Password string
	Host     string
	PORT     int
}

func OpenConnection(config DatabaseConnectionConfig) (*sqlx.DB, error) {
	//const (
	//	host     = "localhost"
	//	port     = 5432
	//	user     = "postgres"
	//	password = "1234"
	//	dbname   = "postgres"
	//)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.PORT, config.User, config.Password, config.DbName)
	return sqlx.Connect("postgres", psqlInfo)
}
