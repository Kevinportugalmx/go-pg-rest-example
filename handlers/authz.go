package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"abc.com/db"
	"abc.com/token"

	"abc.com/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var body models.Login
	if err = json.Unmarshal(bodyBytes, &body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User

	result := db.DB.Where(models.User{Email: body.Email, Password: body.Password}).First(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("INVALID_CREDENTIALS"))
		return
	}

	tokenString, err := token.GenerateToken(models.TokenClaims{
		Email: user.Email,
		Scope: "user",
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	refreshToken, err := token.GenerateRefreshToken(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	expiration := time.Now().Add(24 * time.Hour * 30)
	reponse := models.ResponseToken{
		Token:        tokenString,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Unix(expiration.Unix(), 0),
	}

	jsonResponse, err := json.Marshal(reponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(jsonResponse))
}
