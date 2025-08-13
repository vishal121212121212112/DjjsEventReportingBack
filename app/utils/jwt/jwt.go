// app/utils/jwt/jwt.go
package jwtutil

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

func Sign(sub any, email, secret string) (string, error) {
    claims := jwt.MapClaims{
        "sub":   fmt.Sprint(sub),
        "email": email,
        "iat":   time.Now().Unix(),
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    }
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return t.SignedString([]byte(secret))
}
