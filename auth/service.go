package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string)(*jwt.Token, error)// validasi
}

type jwtService struct {
}

var SECRET_KEY = []byte("anas-startup-secret-key") //jangan sampai orang lain tau, contoh

func NewService() *jwtService{
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user-id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil{
		return signToken, err
	}

	return signToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string)(*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func (token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //hs256
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}