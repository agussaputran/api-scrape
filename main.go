package main

import (
	"scraping-api/config"
	"scraping-api/controllers"
	"scraping-api/models"
)

func main() {
	// db connect to postgre
	pgDB := config.Connect()
	strDB := controllers.StrDB{DB: pgDB}

	// Migrations
	models.Migrations(pgDB)

	strDB.PostDataFromUsersAPI()
	strDB.GetUsersData()

	strDB.PostDataFromJokesAPI()
	strDB.GetJokesData()

	strDB.PostDataFromDogsAPI()
	strDB.GetDogsData()
}
