import { OTPClientConfig, OTPResponse, HealthResponse, TwilioConfig, AWSConfig, LoginRequest, RegisterRequest, LoginResponse, RegisterResponse, MetricsResponse, DLQResponse, QuickSetupConfig } from "./types";
export declare class OTPClient {
    private client;
    private config;
    private providerConfigured;
    constructor(config: OTPClientConfig);
    private ensureProvider;
    quickSetup(config: QuickSetupConfig): Promise<OTPResponse>;
    health(): Promise<HealthResponse>;
    sendOTP(phone: string): Promise<OTPResponse>;
    verifyOTP(phone: string, otp: string): Promise<OTPResponse>;
    addTwilio(config: TwilioConfig): Promise<OTPResponse>;
    addAWS(config: AWSConfig): Promise<OTPResponse>;
    createAPIKey(userId: string): Promise<OTPResponse>;
    register(config: RegisterRequest): Promise<RegisterResponse>;
    login(config: LoginRequest): Promise<LoginResponse>;
    getMetrics(): Promise<MetricsResponse>;
    getDLQ(): Promise<DLQResponse>;
}
