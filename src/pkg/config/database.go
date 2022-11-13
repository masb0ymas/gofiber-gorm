package config

import (
	"fmt"
	"gofiber-gorm/src/database/entity"
	"gofiber-gorm/src/database/seeds"
	"log"
	"strconv"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Declare the variable for the database
var DB *gorm.DB

type QueryFiltered struct {
	Page     string
	PageSize string
	Filtered []string
	Order    []string
}

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
		log.Println("Error database port")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, port, DB_USERNAME, DB_PASSWORD, DB_DATABASE)

	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// List Auto Migrate Table from struct model
	DB.AutoMigrate(&entity.Role{}, &entity.User{}, &entity.Session{})

	// initial seed
	seeds.InitialSeed(DB)

	cyan := color.New(color.FgCyan).SprintFunc()
	dbName := cyan(DB_CONNECTION) + " : " + cyan(DB_DATABASE)

	fmt.Println("Connection " + dbName + " has been established successfully.")
}

// Get Database
func GetDB() *gorm.DB {
	return DB
}
