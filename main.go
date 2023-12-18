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

func validateCreditCardNumber(context *gin.Context) {
	var payload creditCard

	if err := context.BindJSON(&payload); err != nil {
		return
	}

	var r = response{Valid: luhn.Valid(payload.Number)}
	context.IndentedJSON(http.StatusCreated, r)
}

func main() {
	router := gin.Default()
	router.POST("/credit-cards", validateCreditCardNumber)
	router.Run("localhost:8000")
}
