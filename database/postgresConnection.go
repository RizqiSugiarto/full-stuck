package database

import (
	"fmt"
	"list/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionPostgre(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.PostgresDb.Host, cnf.PostgresDb.User, cnf.PostgresDb.Password, cnf.PostgresDb.Dbname, cnf.PostgresDb.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to postgres: %s", err)
	}
	return db
}
