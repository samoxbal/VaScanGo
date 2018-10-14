package controllers

import (
	"VaScanGo/models"
	"VaScanGo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
)

func TokenController(connection *bongo.Connection) iris.Handler {
	return func(ctx iris.Context) {
		var validate *validator.Validate
		validate = validator.New()
		var req models.LoginRequest
		validate.RegisterStructValidation(utils.ValidateLoginStruct)
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.Application().Logger().Infof("Error read request: %s", err)
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		validationError := validate.Struct(req)
		if validationError != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		user := &models.User{}
		userNotFoundError := connection.Collection("users").FindOne(bson.M{
			"name": req.User,
			"password": req.Password,
		}, user)
		if userNotFoundError != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			return
		}
		loginClaims := &models.LoginClaims{
			UserID: user.ID,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, loginClaims)
		tokenString, tokenErr := token.SignedString([]byte("secret"))
		if tokenErr != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			return
		}
		ctx.JSON(map[string]string{
			"token": tokenString,
		})
	}
}
