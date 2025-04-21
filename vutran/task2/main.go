package main

import (
	"log"
	"task2/database"
	"task2/routes"
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

	gin := routes.InitialRouter(db)

	gin.Run()

}

// require (
// 	github.com/gin-gonic/gin v1.10.0
// 	github.com/jmoiron/sqlx v1.4.0
// 	github.com/joho/godotenv v1.5.1
// 	github.com/lib/pq v1.10.9
// )