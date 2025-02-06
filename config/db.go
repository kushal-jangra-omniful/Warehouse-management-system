package config

import (
	"context"
	"fmt"

	// "log"
	// "os"
	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)
var DB *gorm.DB
func InitializeDatabase() {
	Config := postgres.DBConfig{
		Host:                  "localhost",
		Port:                  "5432",
		Username:              "postgres",
		Password:              "kushaljangra",
		Dbname:                "project2",
		DebugMode:             true,
		SkipDefaultTransaction: false,
		PrepareStmt:           true,
		MaxOpenConnections:    100,
		MaxIdleConnections:    10,
		ConnMaxLifetime:       300,
	}

	// Initialize the DB instance
	db := postgres.InitializeDBInstance(Config, &[]postgres.DBConfig{Config})
	DB= db.GetMasterDB(context.Background())
	fmt.Println("✅ Database connected successfully!", DB)
}
// var DB *gorm.DB

// func ConnectDatabase() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	dsn := "host=localhost user=postgres password=kushaljangra dbname=project2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 	}
	
// 	DB = db
// 	fmt.Println("✅ Database connected successfully!")
// }
