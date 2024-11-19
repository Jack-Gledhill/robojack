package utils

import "math/rand"

// Letters is a string containing all ASCII letters, including both cases
const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandString generates a random string of length n, containing only ASCII letters
func RandString(n int) string {
	o := make([]byte, n)
	for i := range o {
		o[i] = Letters[rand.Intn(len(Letters))]
	}

	return string(o)
}
