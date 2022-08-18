package main

import "math/rand"

func generateRandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!_?&/()="
	var randomString = ""
	for i := 0; i < 64; i++ {
		n := rand.Intn(len(charset))
		randomString += string(charset[n])
	}
	return randomString
}
