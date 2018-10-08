package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
)

type GraphQlRequest struct {
	Query string `json:"query"`
}

type GraphQlMutation struct {
	Mutation string `json:"mutation"`
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
	ID 					string 	`json:"id"bson:"_id"`
	User 				string  `json:"user"bson:"user"`
	Name 				string  `json:"name"bson:"name"`
	Description 		string  `json:"description"bson:"description"`
	StartDate 			string  `json:"startDate"bson:"startDate"`
	EndDate 			string  `json:"endDate"bson:"endDate"`
}