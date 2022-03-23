package appier

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RequestBody ...
type RequestBody struct {
	Body []byte `json:"body"`
}

// Payload ...
type Payload struct {
	Data []byte
}

// AppID custom ObjectID
type AppID = primitive.ObjectID

// BrandPayload ...
type BrandPayload struct {
	ID           AppID       `json:"_id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	SearchString string      `json:"searchString"`
	Active       bool        `json:"active"`
	Photos       []string    `json:"photos"`
	Desc         string      `json:"desc"`
	TotalProduct int64       `json:"totalProduct"`
	Logo         string      `json:"logo"`
	Country      CountryInfo `json:"country"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// CountryInfo ...
type CountryInfo struct {
	ID   AppID  `json:"_id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
