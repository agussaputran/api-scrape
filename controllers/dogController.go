package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scraping-api/models"
)

// PostDataFromDogsAPI func route
func (strDB *StrDB) PostDataFromDogsAPI() {
	var dog models.Dogs

	res, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &dog); error != nil {
		fmt.Println("Error ", error.Error())
	}

	strDB.DB.Create(&dog)
}

// GetDogsData func
func (strDB *StrDB) GetDogsData() {
	var dog models.Dogs

	result := strDB.DB.First(&dog)
	fmt.Println(result)
}
