package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func encryptString(text string) string {
	var result string = text
	result += os.Getenv("SQL_SALT")
	fmt.Println(result)
	b := []byte(result)
	h := sha256.New()
	sha := h.Sum(b)
	result = hex.EncodeToString(sha)
	return result
}
