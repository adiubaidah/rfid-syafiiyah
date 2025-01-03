package util

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"strings"
)

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// function string to lower and snake case
func ToSnakeCase(s string) string {
	var result string
	for i, c := range s {
		if i > 0 && 'A' <= c && c <= 'Z' {
			result += "_"
		}
		result += string(c)
	}
	return result
}
func GetDeviceName(topic string) string {
	return strings.Split(topic, "/")[0]
}

func GetDeviceMode(topic string) string {
	return strings.Split(topic, "/")[2]
}

func Generate32ByteKey() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Printf("Failed to generate object key: %v\n", err)
		return ""
	}
	key := hex.EncodeToString(bytes)
	return key
}
