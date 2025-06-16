package shortener

import (
	"math/rand"
	"time"
)

/*
This is the shortener package. Its sole responsibility is to handle the
"business logic" of creating a short code. Keeping it separate makes the logic
reusable and easy to test independently.
*/

const (
	// The character set for generating random short codes.
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// The length of the generated short codes. 6 characters gives 56.8 billion possibilities.
	codeLength = 6
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateShortCode creates a random string of a fixed length from the charset.
func GenerateShortCode() string {
	b := make([]byte, codeLength)
	for i := range b {
		// `seededRand.Intn` returns a random integer in [0,n).
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
