package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Authorization header is required")
			return
		}

		token, err := parseToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "invalid token")
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, err
}

func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetUsernameFromToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) < 2 {
		return ""
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := parseToken(tokenString)
	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["username"].(string); ok {
			return username
		}
	}
	return ""
}
func GetUserIdFromToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) < 2 {
		return ""
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := parseToken(tokenString)
	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["user_id"].(string); ok {
			return username
		}
	}
	return ""
}
