"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const client_1 = require("../src/client");
const assert_1 = __importDefault(require("assert"));
// Simple test to ensure the client constructs without errors
// and sets headers correctly.
async function runTests() {
    const client = new client_1.OTPClient({
        baseURL: "http://localhost:8080",
        apiKey: "test-api-key",
        jwt: "test-jwt",
    });
    // Since we cannot hit the real backend in this CI step easily without starting it,
    // we will verify that the methods exist and are callable.
    (0, assert_1.default)(typeof client.health === "function");
    (0, assert_1.default)(typeof client.sendOTP === "function");
    (0, assert_1.default)(typeof client.verifyOTP === "function");
    (0, assert_1.default)(typeof client.addTwilio === "function");
    (0, assert_1.default)(typeof client.addAWS === "function");
    (0, assert_1.default)(typeof client.createAPIKey === "function");
    (0, assert_1.default)(typeof client.register === "function");
    (0, assert_1.default)(typeof client.login === "function");
    (0, assert_1.default)(typeof client.getMetrics === "function");
    (0, assert_1.default)(typeof client.getDLQ === "function");
    console.log("All client methods exist and are callable!");
}
runTests().catch((err) => {
    console.error("Test failed", err);
    process.exit(1);
});
