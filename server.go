// Package otp provides a framework for multi-tenant OTP generation, verification, and provider management.
package otp

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/bxbx1205/go-otp-framework/config"
	"github.com/bxbx1205/go-otp-framework/models"
	"github.com/bxbx1205/go-otp-framework/routes"
	"github.com/bxbx1205/go-otp-framework/services"
	"github.com/bxbx1205/go-otp-framework/workers"
	"github.com/gin-gonic/gin"
)

// DefaultEmbeddedUserID is the default user ID used when running in embedded library mode.
const DefaultEmbeddedUserID = "000000000000000000000000"

// Server represents the main OTP framework server and embedded client.
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

	initOnce sync.Once
}

// New creates and returns a new Server instance.
func New() *Server {
	return &Server{
		router: gin.Default(),
	}
}

// WithMongo sets the MongoDB URI for the server.
func (s *Server) WithMongo(uri string) *Server {
	s.mongoURI = uri
	return s
}

// WithRedis sets the Redis address for the server.
func (s *Server) WithRedis(addr string) *Server {
	s.redisAddr = addr
	return s
}

// WithTwilio configures the global Twilio provider credentials.
func (s *Server) WithTwilio(sid string, token string, phone string) *Server {
	s.twilioSID = sid
	s.twilioToken = token
	s.twilioPhone = phone
	os.Setenv("TWILIO_PHONE_NUMBER", phone)
	return s
}

// WithAWS configures the global AWS SNS provider credentials.
func (s *Server) WithAWS(accessKey string, secretKey string, region string) *Server {
	s.awsAccessKey = accessKey
	s.awsSecretKey = secretKey
	s.awsRegion = region
	return s
}

// SetupRoutes initializes the HTTP REST API routes on the server's router.
func (s *Server) SetupRoutes() {
	routes.SetupRoutes(s.router)
}

func (s *Server) initialize() error {
	s.initOnce.Do(func() {
		if s.mongoURI != "" {
			config.ConnectMongoDB(s.mongoURI)
		} else {
			log.Println("Warning: MongoURI is not set.")
		}
		if s.redisAddr != "" {
			config.ConnectRedis(s.redisAddr)
		} else {
			log.Println("Warning: RedisAddr is not set.")
		}
		if s.twilioSID != "" {
			config.ConnectTwilio(s.twilioSID, s.twilioToken)
		}
		if s.awsAccessKey != "" {
			config.ConnectAWS(s.awsAccessKey, s.awsSecretKey, s.awsRegion)
		}

		go workers.StartSMSWorker()
	})
	return nil
}

// Start initializes all dependencies and begins listening for HTTP requests on the specified address.
func (s *Server) Start(addr string) error {
	if err := s.initialize(); err != nil {
		return err
	}

	s.SetupRoutes()

	log.Printf("Starting OTP Framework server on %s", addr)
	return s.router.Run(addr)
}

// SendOTP generates and sends an OTP to the given phone number.
func (s *Server) SendOTP(phone string) error {
	if err := s.initialize(); err != nil {
		return err
	}
	return services.SendOTP(DefaultEmbeddedUserID, phone)
}

// VerifyOTP checks if the provided OTP matches the one sent to the phone number.
func (s *Server) VerifyOTP(phone string, otp string) error {
	if err := s.initialize(); err != nil {
		return err
	}
	_, err := services.VerifyOTP(DefaultEmbeddedUserID, phone, otp)
	return err
}

// AddTwilio registers a Twilio provider for the embedded user.
func (s *Server) AddTwilio(sid string, token string, phone string) error {
	if err := s.initialize(); err != nil {
		return err
	}
	req := models.UpsertProviderRequest{
		Provider:    "twilio",
		AccountSID:  sid,
		AuthToken:   token,
		PhoneNumber: phone,
	}
	return services.UpsertProvider(DefaultEmbeddedUserID, req)
}

// AddAWS registers an AWS SNS provider for the embedded user.
func (s *Server) AddAWS(accessKey string, secretKey string, region string) error {
	if err := s.initialize(); err != nil {
		return err
	}
	req := models.UpsertProviderRequest{
		Provider:  "aws",
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
	}
	return services.UpsertProvider(DefaultEmbeddedUserID, req)
}

// CreateAPIKey generates a new programmatic API key for the given user ID.
func (s *Server) CreateAPIKey(userID string) (string, error) {
	if err := s.initialize(); err != nil {
		return "", err
	}
	if userID == "" {
		userID = DefaultEmbeddedUserID
	}
	req := models.CreateAPIKeyRequest{
		Name: "Programmatic Key",
	}
	apiKey, err := services.CreateAPIKey(userID, req)
	if err != nil {
		return "", err
	}
	return apiKey.Key, nil
}

// Health performs a health check on the connected dependencies (MongoDB, Redis).
func (s *Server) Health() error {
	if err := s.initialize(); err != nil {
		return err
	}

	if config.MongoClient != nil {
		err := config.MongoClient.Ping(context.Background(), nil)
		if err != nil {
			return fmt.Errorf("mongodb health check failed: %w", err)
		}
	} else {
		return fmt.Errorf("mongodb is not connected")
	}

	if config.RedisClient != nil {
		_, err := config.RedisClient.Ping(config.Ctx).Result()
		if err != nil {
			return fmt.Errorf("redis health check failed: %w", err)
		}
	} else {
		return fmt.Errorf("redis is not connected")
	}

	return nil
}
