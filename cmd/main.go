package main

import (
	"fmt"
	"os"
	"otp-service/config"
	"otp-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectRedis()

	config.ConnectMongoDB()

	router:=gin.Default()

	routes.SetupRoutes(router)

	port:=os.Getenv("PORT")

	fmt.Println("Server running on port", port)


	
	router.Run(":" + port)
}