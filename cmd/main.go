package main

import (
	"github.com/MohammedSalman999/go-otp-varification/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Initializing config
	app := api.Config{Router : router}

	//router

	app.Routes()

	router.Run(":8000")
}