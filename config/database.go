package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "12345678",
		DBName:   "auth",
	}
	return &dbConfig
}

func DbURL(db *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s", db.User, db.Password, db.Host, db.Port, db.DBName,
	)
}
