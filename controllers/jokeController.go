package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scraping-api/models"
)

// PostDataFromJokesAPI func route
func (strDB *StrDB) PostDataFromJokesAPI() {
	var joke models.Jokes

	res, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &joke); error != nil {
		fmt.Println("Error ", error.Error())
	}

	strDB.DB.Create(&joke)
}

// GetJokesData func
func (strDB *StrDB) GetJokesData() {
	var joke models.Jokes

	result := strDB.DB.First(&joke)
	fmt.Println(result)
}
