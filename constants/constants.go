package constants

const (
	// APIBaseURL is the base URL for the VerifySpeed API
	APIBaseURL = "https://api.verifyspeed.com"

	// ClientIPv4AddressHeaderName is the header name for the client IPv4 address
	ClientIPv4AddressHeaderName = "client-ipv4-address"
)

// Method names
const (
	TelegramMessage = "telegram-message"
	WhatsAppMessage = "whatsapp-message"
	SmsOtp         = "sms-otp"
	WhatsAppOtp    = "whatsapp-otp"
) 