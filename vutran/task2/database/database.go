package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewDb() (*sqlx.DB, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbName)

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