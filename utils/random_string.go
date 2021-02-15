package utils

import "math/rand"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Returns random string of the specified length.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
