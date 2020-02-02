package config

import (
	"github.com/go-pg/pg/v9"
)

// PostgresConfig persists the config for our PostgreSQL database connection
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// GetConnection returns our pg database connection
// usage:
// db := config.GetConnection()
// defer db.Close()
func GetConnection() *pg.DB {
	c := GetPostgresConfig()
	db := pg.Connect(&pg.Options{
		Addr:     c.Host + ":" + c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
		PoolSize: 150,
	})
	return db
}

// GetPostgresConfig returns a PostgresConfig pointer with the correct Postgres Config values
func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "gingoadmin_user",
		Password: "gingoadmin_password",
		Database: "gingoadmin_db",
	}

}
