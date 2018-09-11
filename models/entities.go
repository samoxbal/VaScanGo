package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
)

type GraphQlRequest struct {
	Query string `json:"query"`
}

type LoginRequest struct {
	User string `json:"user"`
	Password string `json:"password"`
}

type User struct {
	bongo.DocumentBase `bson:",inline"`
	Id string `json:"id"`
	Name string `json:"name"`
	Password string `json:"name"`
}

type LoginClaims struct {
	UserId string `json:"user"`
	jwt.StandardClaims
}

type Experiment struct {
	bongo.DocumentBase `bson:",inline"`
	Id string `json:"id"`
	User string `json:"user"`
	Name string `json:"name"`
	Description string `json:"description"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
}
