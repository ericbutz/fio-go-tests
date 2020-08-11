package smoketest

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString returns a random string of length int
func RandomString(length int) string {
	return stringWithCharset(length, LegalCharSet)
}

// RandomAddress returns a random FIO address of <domain>@<address> length int
func RandomAddress(domain string, length int) (address string, err error) {
	if length < (len(domain) + 1) {
		return "", fmt.Errorf("address length must be greater than <domain>@ length")
	}
	return domain + "@" + stringWithCharset(length-1, LegalCharSet), nil
}
