package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/models"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func getWHOISData(url string) (models.WhoIsInfo, error) {
	hostname := url[4:]

	api_url := "https://api.ssllabs.com/api/v3/analyze?host=" + hostname

	// Make a GET request to the API
	response, err := http.Get(api_url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return models.WhoIsInfo{}, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return models.WhoIsInfo{}, err
	}

	// Unmarshal JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return models.WhoIsInfo{}, err
	}

	// Extract protocol from data
	protocol, ok := data["protocol"].(string)
	if !ok {
		return models.WhoIsInfo{}, errors.New("Protocol not found in response")
	}
	//whoisparser part
	host, err := whois.Whois(data["host"].(string))
	result, e := whoisparser.Parse(host)
	var creationDate, expirationDate, registrant, email string
	if e == nil {
		//getting the creation date
		creationDate = result.Domain.CreatedDate
		//getting the expiration date
		expirationDate = result.Domain.ExpirationDate
		//getting the registrant name
		registrant = result.Registrant.Name
		//getting the registrant email address
		email = result.Registrant.Email
	}

	//webscraping the logo
	logo := GetImageLogo(hostname)

	info := models.WhoIsInfo{
		Logo:           logo,
		Protocol:       protocol,
		CreationDate:   creationDate,
		ExpirationDate: expirationDate,
		Registrant:     registrant,
		Email:          email,
	}
	return info, nil
}
