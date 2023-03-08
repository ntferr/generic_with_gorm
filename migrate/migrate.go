package main

import (
	"example/dbgeneric/database"
	"example/dbgeneric/models"
	"fmt"
	"log"
)

func main() {
	log.Println("initiating gorm")
	db := database.InitGorm()

	log.Println("executing gorm AutoMigrate")

	err := db.AutoMigrate(
		models.Car{},
		models.Motorbike{},
	)
	if err != nil {
		fmt.Printf("failed to execute gorm: %e", err)
	}

	log.Println("AutoMigrate executed with success")
}
