package main

import (
	"fmt"
	"log"

	database "github.com/matterinc/PlasmaBlockVerifier/database"
)

func main() {
	db, err := database.OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Hello!")
}
