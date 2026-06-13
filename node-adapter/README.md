# @bxbx/go-otp-framework

A lightweight Node.js/TypeScript adapter for the Go OTP Framework.

## Overview

This package acts as a thin HTTP transport layer around the existing Go OTP Framework backend. It contains no duplicated business logic. The Go backend remains the single source of truth for all integrations.

## Installation

```bash
npm install @bxbx/go-otp-framework
```

## Usage

```typescript
import { OTPClient } from "@bxbx/go-otp-framework";

const otp = new OTPClient({
  baseURL: "http://localhost:8080",
  apiKey: "api_live_xxx", // Optional: Programmatic API key
  jwt: "eyJhbGci...", // Optional: JWT for authenticated user routes
});

async function run() {
  // Check health
  console.log(await otp.health());

  // Send an OTP
  await otp.sendOTP("+919876543210");

  // Verify an OTP
  await otp.verifyOTP("+919876543210", "123456");

  // Add Twilio Provider
  await otp.addTwilio({
    sid: "ACxxx",
    token: "tokenxxx",
    phone: "+1234567890",
  });

  // Add AWS Provider
  await otp.addAWS({
    accessKey: "AKIAxxx",
    secretKey: "secretxxx",
    region: "us-east-1",
  });

  // Create API Key
  await otp.createAPIKey("user_obj_id");

  // Register
  await otp.register({
    email: "test@example.com",
    password: "password123",
  });

  // Login
  await otp.login({
    email: "test@example.com",
    password: "password123",
  });

  // Get Metrics
  console.log(await otp.getMetrics());

  // Get Dead Letter Queue (DLQ)
  console.log(await otp.getDLQ());
}

run();
```

## Authentication Headers

The client will automatically inject `x-api-key: <apiKey>` and `Authorization: Bearer <jwt>` if those respective properties are passed into the `OTPClientConfig`.
