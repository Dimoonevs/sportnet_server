package utils

import (
	"errors"
	"time"

	"github.com/Dimoonevs/SportsApp/auth-service/pkg/data"
	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Usrname string
	Email   string
	Id      int32
}

func (w *JwtWrapper) GenerateToken(user data.User) (signedToken string, err error) {
	claims := &jwtClaims{
		Usrname: user.Username,
		Email:   user.Email,
		Id:      user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}
func (w *JwtWrapper) GetUsernameFromToken(token string) (string, error) {
	claims, err := w.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return claims.Usrname, nil
}
