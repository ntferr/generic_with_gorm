package main

import (
	"example/dbgeneric/database"
	"example/dbgeneric/models"
	"log"
)

func main() {
	db := database.InitGorm()

	dbCar := database.NewDatabase[models.Car](db)

	if err := dbCar.Post(NewCar()); err != nil {
		log.Printf("error: %e", err)
	}

	dbMotorbike := database.NewDatabase[models.Motorbike](db)

	if err := dbMotorbike.Post(NewMotorbike()); err != nil {
		log.Printf("error: %e", err)
	}

}

func NewCar() models.Car {
	return models.Car{
		ID:   "1",
		Name: "Polo",
	}
}

func NewMotorbike() models.Motorbike {
	return models.Motorbike{
		ID:     "1",
		Name:   "CJ-600",
		Wheels: 2,
	}
}
