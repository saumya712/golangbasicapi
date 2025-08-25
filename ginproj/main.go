package main

import (
	"GINPROJ/config"
	"GINPROJ/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.Userroutes(r)

	r.Run(":8080")
}
