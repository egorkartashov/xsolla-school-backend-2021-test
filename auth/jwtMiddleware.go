package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	BearerPrefix              = "Bearer"
	AuthorizationHeader       = "Authorization"
	ServerAuthorizationHeader = "X-SERVER-AUTHORIZATION"
)

func JwtAuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString = ""
		tokenString = parseJwtForServerAuthorization(r)
		if len(tokenString) == 0 {
			tokenString = parseJwtForClientAuthorization(r)
		}

		if len(tokenString) == 0 {
			utils.RespondErrorJson(w, http.StatusUnauthorized, "JWT not found")
			return
		}

		if isJwtValid := validateJwt(tokenString, w); !isJwtValid {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func VendorAuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := parseJwtForServerAuthorization(r)
		if len(tokenString) == 0 {
			utils.RespondErrorJson(w, http.StatusUnauthorized,
				fmt.Sprintf("\"%v\" header not found", AuthorizationHeader))
			return
		}

		if isJwtValid := validateJwt(tokenString, w); !isJwtValid {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func parseJwtForServerAuthorization(r *http.Request) string {
	tokenString := r.Header.Get(ServerAuthorizationHeader)
	return tokenString
}

func parseJwtForClientAuthorization(r *http.Request) string {
	var tokenString = ""
	authorizationHeader := r.Header.Get(AuthorizationHeader)
	if len(authorizationHeader) == 0 {
		return ""
	}

	splitToken := strings.Split(authorizationHeader, BearerPrefix+" ")
	if len(splitToken) == 2 {
		tokenString = splitToken[1]
	}

	return tokenString
}

// Validates JWT token and write errors to response.
// Returns true if token is valid; false otherwise
func validateJwt(tokenString string, w http.ResponseWriter) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("LOGIN_API_KEY")
		log.Print(secretKey)
		return []byte(secretKey), nil
	})

	log.Printf("ValErrors: %v", err)
	if err != nil || !token.Valid {
		if valErr, ok := err.(*jwt.ValidationError); ok {
			var errorMessage = getJwtErrorMessage(valErr.Errors)
			if errorMessage != "" {
				utils.RespondErrorJson(w, http.StatusUnauthorized, errorMessage)
				return false
			}
		}

		utils.RespondErrorJson(w, http.StatusUnauthorized, "Could not handle token")
		return false
	}

	return true
}

func getJwtErrorMessage(validationErrors uint32) string {
	var jwtErrorToMessageMap = map[uint32]string{
		jwt.ValidationErrorMalformed:   "Token is malformed",
		jwt.ValidationErrorExpired:     "Token has expired",
		jwt.ValidationErrorNotValidYet: "Token is not valid yet",
	}

	for jwtError, message := range jwtErrorToMessageMap {
		if validationErrors&jwtError != 0 {
			return message
		}
	}

	return ""
}
