package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/verifyspeed/vs-go/constants"
	"github.com/verifyspeed/vs-go/models"
)

// VerifySpeedClient represents a client for interacting with the VerifySpeed API
type VerifySpeedClient struct {
	httpClient *http.Client
	serverKey  string
}

// NewVerifySpeedClient creates a new instance of VerifySpeedClient
func NewVerifySpeedClient(serverKey string) *VerifySpeedClient {
	return &VerifySpeedClient{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
		serverKey: serverKey,
	}
}

// Initialize initializes the verification process
func (c *VerifySpeedClient) Initialize(clientIPv4Address string) (*models.Initialization, error) {
	req, err := http.NewRequest(http.MethodGet, constants.APIBaseURL+"/v1/verifications/initialize", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set(constants.ClientIPv4AddressHeaderName, clientIPv4Address)
	req.Header.Set("Authorization", c.serverKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to initialize, status code: %d", resp.StatusCode)
	}

	var result models.Initialization
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// CreateVerification creates a new verification
func (c *VerifySpeedClient) CreateVerification(methodName, clientIPv4Address string, language *string) (*models.CreatedVerification, error) {
	payload := struct {
		MethodName string  `json:"methodName"`
		Language   *string `json:"language,omitempty"`
	}{
		MethodName: methodName,
		Language:   language,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, constants.APIBaseURL+"/v1/verifications/create", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(constants.ClientIPv4AddressHeaderName, clientIPv4Address)
	req.Header.Set("Authorization", c.serverKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create verification, status code: %d", resp.StatusCode)
	}

	var result models.CreatedVerification
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// VerifyToken verifies the token and returns the verification result
func (c *VerifySpeedClient) VerifyToken(token string) (*models.VerificationResult, error) {
	req, err := http.NewRequest(http.MethodGet, constants.APIBaseURL+"/v1/verifications/result", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("token", token)
	req.Header.Set("Authorization", c.serverKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify token, status code: %d", resp.StatusCode)
	}

	var result models.VerificationResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
} 