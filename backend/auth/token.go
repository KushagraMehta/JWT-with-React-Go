package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var TOKEN_SECRET = []byte("TOKEN_SECRET")

func CreateToken(user_id uuid.UUID, user User) string {

	authToken := jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id,
		"user_name":  user.Username,
		"email":      user.Email,
		"more_data":  "Blah Blah",
		"exp":        time.Now().Add(time.Hour * 24 * 7).Unix(), //Token expires after 7 Days. Should be short so that if get stolen it only do minimum damage.
	}
	AuthToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, authToken).SignedString(TOKEN_SECRET)
	return fmt.Sprintf("Bearer %v", AuthToken)
}

func TokenValid(coookie *http.Cookie) error {
	tokenString := ExtractToken(coookie)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return TOKEN_SECRET, nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid tokken")
}

func ExtractToken(coookie *http.Cookie) string {
	bearerToken := coookie.Value
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(cookie *http.Cookie) (uuid.UUID, error) {
	tokenString := ExtractToken(cookie)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return TOKEN_SECRET, nil
	})

	if err != nil {
		return uuid.Nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return uuid.MustParse(fmt.Sprintf("%v", claims["user_id"])), nil
	}
	return uuid.Nil, errors.New("internal error")
}

// Hash Funtion do the Hashing of given String.
func Hash(password string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedValue), err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
