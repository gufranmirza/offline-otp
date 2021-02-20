package generator

import (
	"crypto/sha256"
	"fmt"
	"math"
	"time"
)

type generator struct {
	config *Config
}

// New returns new generator implementation
func New(c *Config) OTPGenerater {
	if c.ExpiryTimeInSeconds <= 0 {
		c.ExpiryTimeInSeconds = 30
	}

	if c.OTPLength == 0 {
		c.OTPLength = 6
	}

	if len(c.Secret) == 0 {
		c.Secret = "vjkhvkdfsv8d854vd65f4v65sdf4v65dsf4v"
	}

	return &generator{
		config: c,
	}
}

// Generate Generates OTP with Specified Config
func (g *generator) Generate() int {
	var timeStamp = time.Now().UTC().Unix()
	inputBuffer := fmt.Sprintf("%s%d", g.config.Secret, timeStamp)
	fmt.Println("GenerateOTP Buffer => ", inputBuffer)
	hash := g.calculateHash(inputBuffer)

	return reduceByteArray(hash, g.config)
}

// Validate Validates OTP with Specified Config and timeframe
func (g *generator) Validate(otp int) bool {
	for i := int64(0); i <= g.config.ExpiryTimeInSeconds; i++ {
		var timeStamp = time.Now().UTC().Unix() - i
		inputBuffer := fmt.Sprintf("%s%d", g.config.Secret, timeStamp)

		hash := g.calculateHash(inputBuffer)
		if reduceByteArray(hash, g.config) == otp {
			fmt.Println("ValidateOTP Buffer => ", inputBuffer)
			return true
		}
	}

	return false
}

func (g *generator) calculateHash(inputBuffer string) []byte {
	hash := sha256.New()
	hash.Write([]byte(inputBuffer))
	return hash.Sum(nil)
}

func reduceByteArray(b []byte, c *Config) int {
	asciSummation := int(math.Pow(10, float64(c.OTPLength-1)))
	for _, item := range b {
		asciSummation += int(item)
	}

	return asciSummation
}
