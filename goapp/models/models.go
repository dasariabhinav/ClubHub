package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company Company            `json:"company" bson:"company"`
}

// Company model
type Company struct {
	Owner       Owner       `json:"owner" bson:"owner"`
	Information Info        `json:"information" bson:"information"`
	Franchises  []Franchise `json:"franchises" bson:"franchises"`
}

// Owner model
type Owner struct {
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Contact   Contact `json:"contact" bson:"contact"`
}

// Info model
type Info struct {
	Name      string   `json:"name" bson:"name"`
	TaxNumber string   `json:"tax_number" bson:"tax_number"`
	Location  Location `json:"location" bson:"location"`
}

// Franchise model
type Franchise struct {
	Name      string    `json:"name" bson:"name"`
	URL       string    `json:"url" bson:"url"`
	Location  Location  `json:"location" bson:"location"`
	WhoIsInfo WhoIsInfo `json:"whois,omitempty" bson:"whois,omitempty"`
}

// Location model
type Location struct {
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Address string `json:"address" bson:"address"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
}

// Contact model
type Contact struct {
	Email    string   `json:"email" bson:"email"`
	Phone    string   `json:"phone" bson:"phone"`
	Location Location `json:"location" bson:"location"`
}
type WhoIsInfo struct {
	Logo           string `json:"logo,omitempty" bson:"logo,omitempty"`
	Protocol       string `json:"protocol,omitempty" bson:"protocol,omitempty"`
	CreationDate   string `json:"creation_date,omitempty" bson:"creation_date,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty" bson:"expiration_date,omitempty"`
	Registrant     string `json:"registrant,omitempty" bson:"registrant,omitempty"`
	Email          string `json:"email,omitempty" bson:"email,omitempty"`
}
