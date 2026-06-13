# @bxbx/go-otp-framework

The official Node.js SDK for the Go OTP Framework. Send SMS, verify OTPs, and manage your infrastructure in under 30 seconds.

## Quick Start

The simplest way to get started:

```js
import { OTPClient } from "@bxbx/go-otp-framework";

const otp = new OTPClient({
  baseURL: "http://localhost:8080",
  twilio: {
    sid: "AC...",
    token: "...",
    phone: "+1234567890"
  }
});

await otp.sendOTP("+919876543210");
await otp.verifyOTP("+919876543210", "123456");
```

## Quick Setup (Programmatic)

If you need to configure your provider later in your app's lifecycle:

```ts
// Configure Twilio
await otp.quickSetup({
  provider: "twilio",
  sid: "AC...",
  token: "...",
  phone: "+1234567890"
});

// Configure AWS SNS
await otp.quickSetup({
  provider: "aws",
  accessKey: "AKIA...",
  secretKey: "...",
  region: "ap-south-1"
});
```

---

## Advanced Configuration

For infrastructure management, authentication, and monitoring.

### Setup and Health

```typescript
const otp = new OTPClient({
  baseURL: "http://localhost:8080",
  apiKey: "api_live_xxx", // Optional: Programmatic API key
  jwt: "eyJhbGci...", // Optional: JWT for authenticated user routes
});

// Check health
console.log(await otp.health());
```

### Provider Management (Manual)
Instead of `quickSetup`, you can explicitly add providers:
```typescript
await otp.addTwilio({ sid: "AC...", token: "...", phone: "+1..." });
await otp.addAWS({ accessKey: "AKIA...", secretKey: "...", region: "us-east-1" });
```

### Identity & Access Management
```typescript
// Register User
await otp.register({ email: "test@example.com", password: "password123" });

// Login User
await otp.login({ email: "test@example.com", password: "password123" });

// Create an API Key for automated access
await otp.createAPIKey("user_obj_id");
```

### Monitoring
```typescript
// View aggregate metrics
console.log(await otp.getMetrics());

// Inspect the Dead Letter Queue for failed deliveries
console.log(await otp.getDLQ());
```

## Authentication Headers
The client will automatically inject `x-api-key: <apiKey>` and `Authorization: Bearer <jwt>` if those properties are passed into the `OTPClientConfig`.
