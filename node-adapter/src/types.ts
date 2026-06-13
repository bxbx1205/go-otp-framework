export interface OTPClientConfig {
  baseURL: string;
  apiKey?: string;
  jwt?: string;
}

export interface OTPResponse {
  message: string;
  data?: any;
}

export interface HealthResponse {
  status: string;
  [key: string]: any;
}

export interface TwilioConfig {
  sid: string;
  token: string;
  phone: string;
}

export interface AWSConfig {
  accessKey: string;
  secretKey: string;
  region: string;
}

export interface RegisterRequest {
  email: string;
  password?: string;
}

export interface LoginRequest {
  email: string;
  password?: string;
}

export interface RegisterResponse {
  message: string;
  user: any;
}

export interface LoginResponse {
  message: string;
  token: string;
}

export interface MetricsResponse {
  sms_sent: number;
  sms_failed: number;
  [key: string]: any;
}

export interface DLQResponse {
  jobs: any[];
}
