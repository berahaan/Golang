package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

//  crypto/math is good for security than math/random which is more predictables

func GenerateOTP(length int) string {
	maxNum := big.NewInt(1000000)

	// generate a randowm numbers betwen 0 to 999999
	num, err := rand.Int(rand.Reader, maxNum)
	// incase err happens let us make fall back to time
	if err != nil {
		return fmt.Sprintf("%0*d", length, time.Now().UnixNano()%1000000)
	}
	// other wise return the random number
	return fmt.Sprintf("%0*d", length, num)
}
