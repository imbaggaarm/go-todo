package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

const (
	kTokenPassword string = "TODO_APP_TOKEN_PASSWORD"
)

type Token struct {
	jwt.StandardClaims
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

func CreateTokenString(userID uint, email string) string {
	tk := Token{UserID: userID, Email: email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TODO_APP_TOKEN_PASSWORD")))
	return tokenString
}

func ValidateToken(r *http.Request) (*Token, bool) {
	strToken := ExtractToken(r)
	if strToken == "" {
		return nil, false
	}
	tk := &Token{}
	token, err := jwt.ParseWithClaims(strToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(kTokenPassword)), nil
	})

	if err != nil {
		return nil, false
	}

	return tk, token.Valid
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
