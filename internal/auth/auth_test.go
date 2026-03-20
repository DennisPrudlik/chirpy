package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	gotUserID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}

	if gotUserID != userID {
		t.Fatalf("expected userID %s, got %s", userID, gotUserID)
	}
}

func TestValidateJWTRejectsExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}

func TestValidateJWTRejectsWrongSecret(t *testing.T) {
	userID := uuid.New()
	rightSecret := "right-secret"
	wrongSecret := "wrong-secret"

	token, err := MakeJWT(userID, rightSecret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	_, err = ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Fatal("expected error for wrong secret, got nil")
	}
}

func TestGetBearerToken(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer test-token")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("GetBearerToken returned error: %v", err)
	}

	if token != "test-token" {
		t.Fatalf("expected token test-token, got %s", token)
	}
}

func TestGetBearerTokenMissingHeader(t *testing.T) {
	headers := make(http.Header)

	_, err := GetBearerToken(headers)
	if err == nil {
		t.Fatal("expected error for missing header, got nil")
	}
}

func TestMakeRefreshToken(t *testing.T) {
	token := MakeRefreshToken()
	if token == "" {
		t.Fatal("expected non-empty token")
	}

	if len(token) != 64 {
		t.Fatalf("expected token length 64, got %d", len(token))
	}
}

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey test-api-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey returned error: %v", err)
	}

	if apiKey != "test-api-key" {
		t.Fatalf("expected api key test-api-key, got %s", apiKey)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := make(http.Header)

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error for missing header, got nil")
	}
}
