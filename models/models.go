package models

import "time"

type ResultRequest struct {
	Status  bool        `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type Bill struct {
	Id          int        `json:"id" db:"id"`                   // Bill Id
	Description string     `json:"description" db:"description"` // What you are paying
	ExpDay      int        `json:"expDay" db:"expDay"`           // Expiration day
	Path        string     `json:"path" db:"path"`               // Where to find the files
	LastDate    *time.Time `json:"lastDate" db:"lastDate"`       // Date of last payment
	Track       *bool      `json:"track" db:"track"`             // Is that bill active?
}
