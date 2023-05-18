package config

import (
	"fmt"
	"gofiber-gorm/src/database/migrations"
	"gofiber-gorm/src/database/seeds"
	"gofiber-gorm/src/pkg/helpers"
	"strconv"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	DB_CONNECTION := Env("DB_CONNECTION", "postgres")
	DB_HOST := Env("DB_HOST", "127.0.0.1")
	DB_PORT := Env("DB_PORT", "5432")
	DB_DATABASE := Env("DB_DATABASE", "gofiber-gorm")
	DB_USERNAME := Env("DB_USERNAME", "postgres")
	DB_PASSWORD := Env("DB_PASSWORD", "postgres")

	var err error
	port, err := strconv.ParseUint(DB_PORT, 10, 32)

	if err != nil {
		fmt.Println(err)
		panic("invalid database port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, port, DB_USERNAME, DB_PASSWORD, DB_DATABASE)

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// initial seed
	seeds.Initialize(DB)
	migrations.Initialize(DB)

	cyan := color.New(color.FgCyan).SprintFunc()
	dbName := cyan(DB_CONNECTION) + " : " + cyan(DB_DATABASE)
	msg := fmt.Sprintf("Connection %s has been established successfully.", dbName)

	logMessage := helpers.PrintLog("GORM", msg)
	fmt.Println(logMessage)
}

// Get Database
func GetDB() *gorm.DB {
	return DB
}
