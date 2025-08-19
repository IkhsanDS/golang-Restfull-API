package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "3307")
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASSWORD", "")
	name := getEnv("DB_NAME", "golang")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)
}

func Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(DSN()), &gorm.Config{})
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
