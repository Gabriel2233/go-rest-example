package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mural-app/server/model"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := os.Getenv("dsn")

	db, err = gorm.Open(os.Getenv("dbType"), dsn)

	if err != nil {
		log.Fatal("Error while connection to DB")
	}

	db.LogMode(true)

	db.AutoMigrate(&model.Thought{})

	return db
}
