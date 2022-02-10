package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pi-rate14/golang-jwt-project/database"
	helper "github.com/pi-rate14/golang-jwt-project/helpers"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() 

func VerifyPassword()

func Signup()

func Login()

func GetUser()

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
		err := helper.MatchUserTypeToUid(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

