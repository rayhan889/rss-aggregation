package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	apiKey := strings.Split(authHeader, " ")
	if len(apiKey) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if apiKey[0] != "ApiKey" {
		return "", errors.New("auth header is not ApiKey")
	}

	return apiKey[1], nil
}