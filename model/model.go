package model

import (
	"fmt"

	"github.com/invoicepro360/go-common/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes database connection
func Connect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatalf("Failed to create database connection: %s", err.Error())
	}

	return db
}
