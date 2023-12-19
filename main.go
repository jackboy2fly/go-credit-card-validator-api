package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theplant/luhn"
)

type creditCard struct {
	Number int `json:"number"`
}

type response struct {
	Valid bool `json:"isValid"`
}

func postHandler(context *gin.Context) {
	var payload creditCard

	if err := context.BindJSON(&payload); err != nil {
		return
	}

	var r = response{Valid: luhn.Valid(payload.Number)}
	context.JSON(http.StatusCreated, r)
}

func getHandler(context *gin.Context) {
	m := map[string]string{"message": "Welcome to this credit card validator API!"}
	context.IndentedJSON(http.StatusOK, m)
}

func main() {
	router := gin.Default()
	router.GET("/", getHandler)
	router.POST("/credit-cards", postHandler)
	router.Run(":8000")
}
