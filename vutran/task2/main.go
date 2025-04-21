package main

import (
	"fmt"
	"log"
	"task2/database"
	"task2/routes"
)

func main() {
	db, err := database.NewDb()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	err = database.InitSchema(db)

	if err != nil {
		log.Fatal(err)
	}

	gin := routes.InitialRouter(db)

	gin.Run()

}
