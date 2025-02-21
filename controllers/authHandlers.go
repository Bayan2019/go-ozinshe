package controllers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Bayan2019/go-ozinshe/repositories"
	"github.com/Bayan2019/go-ozinshe/views"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	// TokenTypeAccess -
	// Set the Issuer to "chirpy"
	TokenTypeAccess TokenType = "chirpy-access"
)

type AuthHandlers struct {
	authRepo *repositories.UsersRepository
}

func NewAuthHandlers(repo *repositories.UsersRepository) *AuthHandlers {
	return &AuthHandlers{
		authRepo: repo,
	}
}

func (ah *AuthHandlers) MiddlewareAuth(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := GetBearerToken(r.Header)
		if err != nil {
			views.RespondWithError(w, http.StatusUnauthorized, "Couldn't find token", err)
			return
		}

		email, err := ValidateJWT(jwtToken, string(TokenTypeAccess))

		_, err = ah.authRepo.DB.GetUserByEmail(r.Context(), email)
		if err != nil {
			views.RespondWithError(w, http.StatusNotFound, "Couldn't get user", err)
			return
		}

		handler.ServeHTTP(w, r)
	}
}

// 6. Authentication / 1. Authentication with JWTs
// Add a GetBearerToken function to your auth package
// GetBearerToken -
func GetBearerToken(headers http.Header) (string, error) {
	// Auth information will come into our server
	// in the Authorization header.
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		// If the header doesn't exist, return an error.
		return "", errors.New("no auth header included in request")
	}
	// stripping off the Bearer prefix and whitespace
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "Bearer" {
		// If the header doesn't exist, return an error.
		return "", errors.New("malformed authorization header")
	}
	// return the TOKEN_STRING if it exists
	return splitAuth[1], nil
}

// 6. Authentication / 6. JWTs
// Add a MakeJWT function to your auth package:
// MakeJWT -
func MakeJWT(
	email string,
	tokenSecret string,
	expiresIn time.Duration,
) (string, error) {
	signingKey := []byte(tokenSecret)
	// Use jwt.NewWithClaims to create a new token
	token := jwt.NewWithClaims(
		// Use jwt.SigningMethodHS256 as the signing method.
		jwt.SigningMethodHS256,
		// Use jwt.RegisteredClaims as the claims
		jwt.RegisteredClaims{
			Issuer: string(TokenTypeAccess),
			// Set IssuedAt to the current time in UTC
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
			// Set ExpiresAt to the current time plus the expiration time (expiresIn)
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
			// Set the Subject to a stringified version of the user's id
			Subject: email,
		})
	// Use token.SignedString to sign the token with the secret key.
	return token.SignedString(signingKey)
}

// 6. Authentication / 6. JWTs
// Add a ValidateJWT function to your auth package:
// ValidateJWT -
func ValidateJWT(tokenString, tokenSecret string) (string, error) {
	claimsStruct := jwt.RegisteredClaims{}
	// Use the jwt.ParseWithClaims function
	// to validate the signature of the JWT and extract the claims
	// into a *jwt.Token struct.
	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenSecret), nil
		},
	)
	if err != nil {
		// An error will be returned if the token is invalid or has expired.
		// If the token is invalid,
		// return a 401 Unauthorized response from your handler.
		return "", err
	}

	// If all is well with the token,
	// use the token.Claims interface
	// to get access to the user's id from the claims
	// (which should be stored in the Subject field).
	email, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", err
	}
	if issuer != string(TokenTypeAccess) {
		return "", errors.New("invalid issuer")
	}

	return email, nil
}
