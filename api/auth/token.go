package auth

import (
	"github.com/dgrijalva/jwt-go"
	"go-todo/api/model"
	"net/http"
	"os"
	"strings"
)

const (
	kTokenPassword string = "TODO_APP_TOKEN_PASSWORD"
)

func ValidateToken(r *http.Request) bool {
	strToken := ExtractToken(r)
	if strToken == "" {
		return false
	}
	tk := &model.Token{}
	token, err := jwt.ParseWithClaims(strToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(kTokenPassword)), nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}
	return true
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	splittedToken := strings.Split(bearerToken, " ")
	if len(splittedToken) == 2 {
		return splittedToken[1]
	}
	return ""
}
