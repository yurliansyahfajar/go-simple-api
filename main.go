package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yurliansyahfajar/go-simple-api/db"
	"github.com/yurliansyahfajar/go-simple-api/routes"
)

func main() {

	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
