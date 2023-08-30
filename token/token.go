package token

import (
	"errors"
	"time"

	"abc.com/models"
	"github.com/golang-jwt/jwt/v4"
)

var mySeed = []byte("som3th1ng_s3cr3t") //.ENV

func GenerateFromRefresh(accesstoken string) (newToken string, newRefreshToken string, err error) {
	claims := models.TokenClaims{}
	token, err := jwt.Parse(accesstoken, func(token *jwt.Token) (interface{}, error) {
		return mySeed, nil
	})

	if err != nil {
		return
	}
	if claimsToken, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.Active = true
		claims.Scope = claimsToken["scope"].(string)
		claims.Email = claimsToken["email"].(string)
	} else {
		err = errors.New("Unauthorized")
		return
	}

	if newToken, err = GenerateToken(claims); err != nil {
		return
	}
	if newRefreshToken, err = GenerateRefreshToken(newToken); err != nil {
		return
	}
	return
}

func GenerateToken(claims models.TokenClaims) (string, error) {
	claims.Active = true
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySeed)
}

func GenerateRefreshToken(accesstoken string) (string, error) {
	claims := &models.TokenClaims{}

	token, err := jwt.Parse(accesstoken, func(token *jwt.Token) (interface{}, error) {
		return mySeed, nil
	})

	if err != nil {
		return "", err
	}
	if claimsToken, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims.Active = true
		claims.Scope = claimsToken["scope"].(string)
		claims.Email = claimsToken["email"].(string)
		claims.RegisteredClaims = jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 30)),
		}
	} else {
		return "", errors.New("Unauthorized")
	}
	refreshAccesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshAccesToken.SignedString(mySeed)
}

func GenerateIntrospection(accesToken string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(accesToken, func(token *jwt.Token) (interface{}, error) {
		return mySeed, nil
	})

	if err != nil {
		return false, nil, err
	}

	claimsToken, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, nil, errors.New("Unauthorized")
	}

	exp := claimsToken["exp"].(float64)
	if expTime := time.Unix(int64(exp), 0); expTime.Before(time.Now()) {
		return false, nil, nil
	}

	return true, claimsToken, nil
}
