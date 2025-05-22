package main

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const length = 10

func CreateRandomString() string {
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[(rand.Intn(len(charset)))]
	}
	return string(b)
}
