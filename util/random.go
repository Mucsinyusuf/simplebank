package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Local().UnixNano())

}

// generates random integers
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generates random strings
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

// generate random owner
func RandomOwner() string {
	return RandomString(6)

}

//generate random money

func RandomMoney() int64 {
	return RandomInt(0, 3000)
}

// generate randomcurrency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}
