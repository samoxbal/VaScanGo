package models

import "github.com/go-bongo/bongo"

type Experiment struct {
	bongo.DocumentBase `bson:",inline"`
	Id string `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
	Description string `json:"description"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
}
