package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const issuer = "test"

var jwtCfg JWT

type JWT struct {
	atSecretKey []byte        //Access Token Secret Key
	atd         time.Duration //Access Token Duration
	rtSecretKey []byte        //Refresh Token Secret Key
	rtd         time.Duration //Refresh Token Duration
}

type Claims struct {
	jwt.StandardClaims
	Session string `json:"session"`
	Renew   string `json:"renew,omitempty"`
}

func CheckAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return jwtCfg.atSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	isr := fmt.Sprintf("%v", claims["iss"])
	if isr != issuer {
		return nil, fmt.Errorf("Invalid Issuer")
	}

	return claims, nil
}
