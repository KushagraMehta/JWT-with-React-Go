package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/KushagraMehta/blog/JWT-with-React+Go/Code/backend/auth"
	"github.com/KushagraMehta/blog/JWT-with-React+Go/Code/backend/middleware"
	"github.com/google/uuid"
)

type Handler struct {
	UserData map[uuid.UUID]auth.User
}

func (h *Handler) Init() {
	h.UserData = make(map[uuid.UUID]auth.User)
}
func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tmpUser auth.User
	json.Unmarshal(reqBody, &tmpUser)
	log.Printf("%+v", tmpUser)
	tmpUser.Password, _ = auth.Hash(tmpUser.Password)
	userID := uuid.New()
	h.UserData[userID] = tmpUser
	generatedTokens := auth.CreateToken(userID, tmpUser)

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    generatedTokens,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	middleware.JSON(w, http.StatusCreated, "OK")
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	middleware.JSON(w, http.StatusOK, "OK")
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	middleware.JSON(w, http.StatusOK, "OK")
}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, _ := r.Cookie("jwt")
	uid, err := auth.ExtractTokenID(tokenCookie)
	if err != nil {
		fmt.Println(err)
	}
	userData := h.UserData[uid]
	middleware.JSON(w, http.StatusOK, map[string]string{
		"username": userData.Username,
		"email":    userData.Email,
	})
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tmpUser auth.User
	json.Unmarshal(reqBody, &tmpUser)
	log.Printf("%+v", tmpUser)

	for uid, value := range h.UserData {
		if (value.Email == tmpUser.Email || value.Username == tmpUser.Email) && auth.VerifyPassword(value.Password, tmpUser.Password) == nil {

			generatedTokens := auth.CreateToken(uid, value)
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    generatedTokens,
				Expires:  time.Now().Add(time.Hour * 24 * 7),
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
				Secure:   true,
			})
			middleware.JSON(w, http.StatusCreated, "OK")
			return
		}
	}
	middleware.JSON(w, http.StatusUnauthorized, "")
}
