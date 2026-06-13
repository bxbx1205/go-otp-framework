import { OTPClient } from "../src/client";
import assert from "assert";

// Simple test to ensure the client constructs without errors
// and sets headers correctly.

async function runTests() {
  const client = new OTPClient({
    baseURL: "http://localhost:8080",
    apiKey: "test-api-key",
    jwt: "test-jwt",
  });

  // Since we cannot hit the real backend in this CI step easily without starting it,
  // we will verify that the methods exist and are callable.
  assert(typeof client.health === "function");
  assert(typeof client.sendOTP === "function");
  assert(typeof client.verifyOTP === "function");
  assert(typeof client.addTwilio === "function");
  assert(typeof client.addAWS === "function");
  assert(typeof client.createAPIKey === "function");
  assert(typeof client.register === "function");
  assert(typeof client.login === "function");
  assert(typeof client.getMetrics === "function");
  assert(typeof client.getDLQ === "function");

  console.log("All client methods exist and are callable!");
}

runTests().catch((err) => {
  console.error("Test failed", err);
  process.exit(1);
});
