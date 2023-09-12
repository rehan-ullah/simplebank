package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// @ Generate Random String
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currencies := []string{"PKR", "AED", "USD"}
	k := len(currencies)
	return currencies[rand.Intn(k)]
}
