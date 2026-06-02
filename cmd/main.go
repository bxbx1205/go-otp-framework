package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"otp-service/config"
	"otp-service/middleware"
	"otp-service/routes"
	"syscall"
	"time"
	"otp-service/workers"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	config.ConnectRedis()

	config.ConnectMongoDB()

	router:=gin.Default()

	router.Use(middleware.LoggerMiddleware())

	routes.SetupRoutes(router)

	port:=os.Getenv("PORT")

	config.ConnectAWS()

	config.ConnectTwilio()

	router.Use(
	middleware.RequestIDMiddleware(),

)

	go workers.StartSMSWorker()

	server := &http.Server{

		Addr: ":" + port,

		Handler: router,
	}

	go func ()  {
		fmt.Println("serving running on port ", port)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal,1)

	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM,)

	<-quit
	fmt.Println("shutting down server")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		log.Fatal(
			"Server forced to shutdown:",
			err,
		)
	}


	fmt.Println("Server exited properly")
	

	// fmt.Println("Server running on port", port)
}