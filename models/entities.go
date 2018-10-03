package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
)

type GraphQlRequest struct {
	Query string `json:"query"`
}

type LoginRequest struct {
	User 		string `json:"user"`
	Password 	string `json:"password"`
}

type User struct {
	ID 					string `json:"id"`
	Name 				string `json:"name"`
	Password 			string `json:"name"`
	bongo.DocumentBase 		   `bson:",inline"`
}

type LoginClaims struct {
	UserID string `json:"user"`
	jwt.StandardClaims
}

type Experiment struct {
	ID 					string `json:"id"`
	User 				string `json:"user"`
	Name 				string `json:"name"`
	Description 		string `json:"description"`
	StartDate 			string `json:"startDate"`
	EndDate 			string `json:"endDate"`
	bongo.DocumentBase 		   `bson:",inline"`
}

type Event struct {
	ID            string      `json:"id"`
	AggregateID   string      `json:"aggregate_id"`
	Type          string      `json:"type"`
	Data          interface{} `json:"data"`
	bongo.DocumentBase 		  `bson:",inline"`
}

type BaseAggregate struct {
	ID      string
	Type    string
	Changes []Event
}

type AggregateHandler interface {
	GetID() string
	ApplyChange(event Event)
	ApplyChangeHelper(aggregate AggregateHandler, event Event)
}

type ExperimentAggregate struct {
	BaseAggregate
	User string
}
