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
	ID           AppID
	Name         string
	Slug         string
	SearchString string
	Active       bool
	Photos       []string
	Desc         string
	TotalProduct int64
	Logo         string
	Country      CountryInfo
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CountryInfo ...
type CountryInfo struct {
	ID   AppID
	Name string
	Code string
}
