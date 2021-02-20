package generator

// Config represents the config accepted by the generator
type Config struct {
	OTPLength           int    `json:"otp_length"`
	ExpiryTimeInSeconds int64  `json:"expiry_in_seconds"`
	Secret              string `json:"secret"`
}

// OTPGenerater interface
type OTPGenerater interface {
	Generate() int
	Validate(otp int) bool
}
