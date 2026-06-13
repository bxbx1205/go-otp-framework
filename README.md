# OTP Framework

A production-grade, reusable Go framework for building Multi-Tenant OTP SaaS Platforms with dependency injection, dynamic provider routing, and background retry queues.

## Overview

This framework allows you to easily embed a scalable OTP generation/verification engine with out-of-the-box support for:

- **Dual Modes**: Run as an embedded library directly calling Go methods, or as a RESTful HTTP microservice.
- **Multi-Tenant Credentials**: Each of your SaaS customers can use their own Twilio or AWS SNS credentials dynamically.
- **Provider Failover**: Easily register Twilio and AWS SNS.
- **Encryption**: AES-256 GCM encryption protects stored provider credentials.
- **Queues & DLQ**: Background SMS job queue with retry counts, exponential fallbacks, and DLQ tracking built natively over Redis.
- **JWT & Authentication**: User management, auth login, and api-key generation middleware.

## Installation

```sh
go get github.com/bxbx1205/go-otp-framework@v1.6.6
```

## Setup Dependencies (Docker)

Before running the server or executing embedded Go code, you must have MongoDB and Redis instances running. The easiest way to get started is by using the provided `docker-compose.yml`:

```sh
# Copy the environment template
cp .env.example .env

# Spin up MongoDB, Redis, and the OTP framework server
docker-compose up -d
```

Once Docker is running, the services will be available at:
- **MongoDB**: `mongodb://admin:password@localhost:27017`
- **Redis**: `localhost:6379`
- **REST API** (if started via Docker): `http://localhost:8080`

## Quick Start Example

You do not need to clone the repository or set environment variables manually anymore. Use dependency injection to initialize the framework.

### A) Embedded Library Mode

Use the framework directly in your Go code without starting an HTTP server. It will automatically connect to databases and start background workers on the first method call.

```go
package main

import (
	"log"

	"github.com/bxbx1205/go-otp-framework"
)

func main() {
	// Initialize framework using the Builder pattern
	server := otp.New().
		WithMongo("mongodb://admin:password@localhost:27017").
		WithRedis("localhost:6379").
		WithTwilio("AC_xxxx", "AUTH_TOKEN", "+12345678").
		WithAWS("AWS_KEY", "AWS_SECRET", "us-east-1")

	// Directly send OTP (lazily initializes dependencies!)
	err := server.SendOTP("+919876543210")
	if err != nil {
		log.Fatal(err)
	}
    
	// Verify OTP
	err = server.VerifyOTP("+919876543210", "123456")
	if err != nil {
		log.Fatal(err)
	}
}
```

### B) REST API Mode

Start a standalone REST API microservice.

```go
package main

import (
	"log"

	"github.com/bxbx1205/go-otp-framework"
)

func main() {
	// Initialize framework
	server := otp.New().
		WithMongo("mongodb://admin:password@localhost:27017").
		WithRedis("localhost:6379")

	// Start the Server on :8080 (also initializes dependencies)
	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Creating API Keys

If using REST API Mode:

1. Register an admin/tenant account via `POST /api/v1/auth/register` and `POST /api/v1/auth/login`.
2. Grab the JWT `token`, and create an API Key `POST /api/v1/apikey/create` utilizing `Authorization: Bearer <token>`.
3. Use the newly issued `x-api-key: api_live_xx...` to issue OTP generation events `POST /api/v1/otp/send`.

If using Embedded Library Mode:
```go
apiKey, err := server.CreateAPIKey("your_user_id")
```

## Dynamic Provider Execution (SaaS Architecture)

Instead of forcing all API traffic through a single global Twilio account, your customers can define their own logic.

Via REST:
`POST /api/v1/providers/twilio`
```json
{
    "sid": "ACyour_customer_sid",
    "token": "their_secret_token",
    "phone": "+123..."
}
```

`POST /api/v1/providers/aws`
```json
{
    "accessKey": "AKIA...",
    "secretKey": "their_secret_key",
    "region": "us-east-1"
}
```

Via Embedded Library:
```go
err := server.AddTwilio("ACyour_customer_sid", "their_secret_token", "+123...")
```

The underlying Worker Queue instantly encrypts this Token via AES-256 for MongoDB storage, and decrypts it into memory only during specific execution windows!
