package utils

import (
	"VaScanGo/models"
	"gopkg.in/go-playground/validator.v9"
)

func ValidateQueryStruct(sl validator.StructLevel) {
	query := sl.Current().Interface().(models.GraphQlRequest)
	if len(query.Query) == 0 {
		sl.ReportError(query.Query, "Query", "query", "", "")
	}
}

func ValidateLoginStruct(sl validator.StructLevel) {
	login := sl.Current().Interface().(models.LoginRequest)
	if len(login.User) == 0 {
		sl.ReportError(login.User, "User", "user", "", "")
	}
	if len(login.Password) == 0 {
		sl.ReportError(login.Password, "Password", "password", "", "")
	}
}