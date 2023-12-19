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

func handler(context *gin.Context) {
	var payload creditCard

	if err := context.BindJSON(&payload); err != nil {
		return
	}

	var r = response{Valid: luhn.Valid(payload.Number)}
	context.JSON(http.StatusCreated, r)
}

func main() {
	router := gin.Default()
	router.POST("/credit-cards", handler)
	router.Run(":8000")
}
