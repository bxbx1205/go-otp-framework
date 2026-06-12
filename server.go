package framework

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/myusername/otp-framework/config"
	"github.com/myusername/otp-framework/routes"
	"github.com/myusername/otp-framework/workers"
)

type Server struct {
	router *gin.Engine

	mongoURI  string
	redisAddr string

	twilioSID   string
	twilioToken string
	twilioPhone string

	awsAccessKey string
	awsSecretKey string
	awsRegion    string
}

func New() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) WithMongo(uri string) *Server {
	s.mongoURI = uri
	return s
}

func (s *Server) WithRedis(addr string) *Server {
	s.redisAddr = addr
	return s
}

func (s *Server) WithTwilio(sid string, token string, phone string) *Server {
	s.twilioSID = sid
	s.twilioToken = token
	s.twilioPhone = phone
	os.Setenv("TWILIO_PHONE_NUMBER", phone) 
	return s
}

func (s *Server) WithAWS(accessKey string, secretKey string, region string) *Server {
	s.awsAccessKey = accessKey
	s.awsSecretKey = secretKey
	s.awsRegion = region
	return s
}

func (s *Server) SetupRoutes() {
	routes.SetupRoutes(s.router)
}

func (s *Server) Start(addr string) error {
	if s.mongoURI != "" {
		config.ConnectMongoDB(s.mongoURI)
	}
	if s.redisAddr != "" {
		config.ConnectRedis(s.redisAddr)
	}
	if s.twilioSID != "" {
		config.ConnectTwilio(s.twilioSID, s.twilioToken)
	}
	if s.awsAccessKey != "" {
		config.ConnectAWS(s.awsAccessKey, s.awsSecretKey, s.awsRegion)
	}

	s.SetupRoutes()

	go workers.StartSMSWorker()

	log.Printf("Starting OTP Framework server on %s", addr)
	return s.router.Run(addr)
}
