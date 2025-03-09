package db

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connection() *goqu.Database {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	dbCtx, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("Connected to database:Error")
		log.Fatal(err)
		os.Exit(-1)
	}
	err = dbCtx.Ping()
	if err != nil {
		fmt.Println("Connected to database:Error")
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println("Connected to database:Success")
	dialect := goqu.Dialect("postgres")
	return dialect.DB(dbCtx)
}

func ConnectionGorm() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connected to database:Error")
	}
	fmt.Println("Connected to database:Success")
	return db
}
