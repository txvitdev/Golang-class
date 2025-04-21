package database

import (
	"fmt"
	"os"
	"task2/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDb() (*sqlx.DB, error) {
	config := config.SetupConfig()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DbName)

	db, err := sqlx.Connect("postgres", connStr)

	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}

func InitSchema(db *sqlx.DB) error {
	schemaBytes, err := os.ReadFile("database/schema.sql")

	if err != nil {
		return err
	} else {
		fmt.Println("Connect to database successfully")
	}

	schema := string(schemaBytes)

	_, err = db.Exec(schema)

	if err != nil {
		return err
	}

	fmt.Print("Create schema successfully!!!")
	return nil
}
