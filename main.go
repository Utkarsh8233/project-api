package main

import (
	"github.com/engnrutkarsh/project-api/db"
	"github.com/engnrutkarsh/project-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8090")
	if err != nil {
		return
	}
}
