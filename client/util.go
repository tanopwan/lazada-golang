package client

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

func randomUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func hmacSHA256(payload string, appSecret string) string {
	secret := []byte(appSecret)
	bb := []byte(payload)
	hash := hmac.New(sha256.New, secret)
	hash.Write(bb)

	sum := hash.Sum(nil)
	return strings.ToUpper(hex.EncodeToString([]byte(sum)))
}

func sign(data url.Values, urlName string, appSecret string) string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	sort.StringSlice(keys).Sort()

	var payload string
	for _, key := range keys {
		payload = payload + key + data.Get(key)
	}

	return hmacSHA256(urlName+payload, appSecret)
}
