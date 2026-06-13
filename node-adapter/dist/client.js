"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.OTPClient = void 0;
const axios_1 = __importDefault(require("axios"));
class OTPClient {
    constructor(config) {
        const headers = {
            "Content-Type": "application/json",
        };
        if (config.apiKey) {
            headers["x-api-key"] = config.apiKey;
        }
        if (config.jwt) {
            headers["Authorization"] = `Bearer ${config.jwt}`;
        }
        this.client = axios_1.default.create({
            baseURL: config.baseURL,
            headers,
        });
    }
    async health() {
        const response = await this.client.get("/health");
        return response.data;
    }
    async sendOTP(phone) {
        const response = await this.client.post("/api/v1/otp/send", { phone });
        return response.data;
    }
    async verifyOTP(phone, otp) {
        const response = await this.client.post("/api/v1/otp/verify", {
            phone,
            otp,
        });
        return response.data;
    }
    async addTwilio(config) {
        const response = await this.client.post("/api/v1/providers/twilio", config);
        return response.data;
    }
    async addAWS(config) {
        const response = await this.client.post("/api/v1/providers/aws", config);
        return response.data;
    }
    async createAPIKey(userId) {
        const response = await this.client.post("/api/v1/apikey/create", { userId });
        return response.data;
    }
    async register(config) {
        const response = await this.client.post("/api/v1/auth/register", config);
        return response.data;
    }
    async login(config) {
        const response = await this.client.post("/api/v1/auth/login", config);
        return response.data;
    }
    async getMetrics() {
        const response = await this.client.get("/api/v1/metrics");
        return response.data;
    }
    async getDLQ() {
        const response = await this.client.get("/api/v1/dlq");
        return response.data;
    }
}
exports.OTPClient = OTPClient;
