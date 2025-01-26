package models

import "time"

// Method represents a verification method
type Method struct {
	MethodName  string `json:"methodName"`
	DisplayName string `json:"displayName"`
}

// Initialization represents the initialization response
type Initialization struct {
	AvailableMethods []Method `json:"availableMethods"`
}

// CreatedVerification represents the result of a created verification request
type CreatedVerification struct {
	MethodName      string `json:"methodName"`
	VerificationKey string `json:"verificationKey"`
	DeepLink        string `json:"deepLink,omitempty"`
}

// VerificationResult represents the result of a verification process
type VerificationResult struct {
	MethodName         string    `json:"methodName"`
	DateOfVerification time.Time `json:"dateOfVerification"`
	PhoneNumber        string    `json:"phoneNumber"`
} 