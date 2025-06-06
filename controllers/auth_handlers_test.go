package controllers

import (
	"net/http"
	"testing"
	"time"
	// "github.com/google/uuid"
)

func TestCheckPasswordHash(t *testing.T) {
	// First, we need to create some hashed passwords for testing
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := hashPassword(password1)
	hash2, _ := hashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 6. Authentication / 6. JWTs
func TestValidateJWT(t *testing.T) {
	userEmail := "admin@admin.com"
	validToken, _ := makeJWT(userEmail, "secret", time.Hour)

	tests := []struct {
		name          string
		tokenString   string
		tokenSecret   string
		wantUserEmail string
		wantErr       bool
	}{
		{
			name:          "Valid token",
			tokenString:   validToken,
			tokenSecret:   "secret",
			wantUserEmail: userEmail,
			wantErr:       false,
		},
		{
			name:          "Invalid token",
			tokenString:   "invalid.token.string",
			tokenSecret:   "secret",
			wantUserEmail: "",
			wantErr:       true,
		},
		{
			name:          "Wrong secret",
			tokenString:   validToken,
			tokenSecret:   "wrong_secret",
			wantUserEmail: "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserEmail, err := validateJWT(tt.tokenString, tt.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserEmail != tt.wantUserEmail {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserEmail, tt.wantUserEmail)
			}
		})
	}
}

// 6. Authentication / 7. Authentication With JWTs
func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "Valid Bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			wantToken: "valid_token",
			wantErr:   false,
		},
		{
			name:      "Missing Authorization header",
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InvalidBearer token"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := getBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}
