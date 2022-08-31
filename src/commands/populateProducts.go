package main

import (
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/maxts0gt/go-ambassador/src/database"
	"github.com/maxts0gt/go-ambassador/src/models"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Username(),
			Description: faker.Username(),
			Image:       faker.URL(),
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&product)
	}
}
