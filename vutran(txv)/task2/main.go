package main

import (
	"context"
	"fmt"
	"log"
	"task2/database"
	"time"
)

func main() {
	db, err := database.NewDb()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	err = database.InitSchema(db)

	if  err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	start := time.Now()

	var result string
	err = db.QueryRowContext(ctx, "SELECT pg_sleep(10); SELECT 'done'").Scan(&result)

	duration := time.Since(start)
	fmt.Println("QueryRowContext result:", result)
	fmt.Println("Error:", err)
	fmt.Println("Duration:", duration)
}