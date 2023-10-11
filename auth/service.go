package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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