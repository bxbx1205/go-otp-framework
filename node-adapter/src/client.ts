import axios, { AxiosInstance } from "axios";
import {
  OTPClientConfig,
  OTPResponse,
  HealthResponse,
  TwilioConfig,
  AWSConfig,
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  RegisterResponse,
  MetricsResponse,
  DLQResponse,
} from "./types";

export class OTPClient {
  private client: AxiosInstance;

  constructor(config: OTPClientConfig) {
    const headers: Record<string, string> = {
      "Content-Type": "application/json",
    };

    if (config.apiKey) {
      headers["x-api-key"] = config.apiKey;
    }

    if (config.jwt) {
      headers["Authorization"] = `Bearer ${config.jwt}`;
    }

    this.client = axios.create({
      baseURL: config.baseURL,
      headers,
    });
  }

  public async health(): Promise<HealthResponse> {
    const response = await this.client.get("/health");
    return response.data;
  }

  public async sendOTP(phone: string): Promise<OTPResponse> {
    const response = await this.client.post("/api/v1/otp/send", { phone });
    return response.data;
  }

  public async verifyOTP(phone: string, otp: string): Promise<OTPResponse> {
    const response = await this.client.post("/api/v1/otp/verify", {
      phone,
      otp,
    });
    return response.data;
  }

  public async addTwilio(config: TwilioConfig): Promise<OTPResponse> {
    const response = await this.client.post("/api/v1/providers/twilio", config);
    return response.data;
  }

  public async addAWS(config: AWSConfig): Promise<OTPResponse> {
    const response = await this.client.post("/api/v1/providers/aws", config);
    return response.data;
  }

  public async createAPIKey(userId: string): Promise<OTPResponse> {
    const response = await this.client.post("/api/v1/apikey/create", { userId });
    return response.data;
  }

  public async register(config: RegisterRequest): Promise<RegisterResponse> {
    const response = await this.client.post("/api/v1/auth/register", config);
    return response.data;
  }

  public async login(config: LoginRequest): Promise<LoginResponse> {
    const response = await this.client.post("/api/v1/auth/login", config);
    return response.data;
  }

  public async getMetrics(): Promise<MetricsResponse> {
    const response = await this.client.get("/api/v1/metrics");
    return response.data;
  }

  public async getDLQ(): Promise<DLQResponse> {
    const response = await this.client.get("/api/v1/dlq");
    return response.data;
  }
}
