package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scraping-api/models"
)

// PostDataFromUsersAPI func route
func (strDB *StrDB) PostDataFromUsersAPI() {
	var user models.Users

	res, err := http.Get("http://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)

	if error := json.Unmarshal([]byte(stringBody), &user); error != nil {
		fmt.Println("Error ", error.Error())
	}

	strDB.DB.Create(&user)
}

// GetUsersData func
func (strDB *StrDB) GetUsersData() {
	var user models.Users

	result := strDB.DB.First(&user)
	fmt.Println(result)
}
