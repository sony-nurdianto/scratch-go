package model

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//Users Model
type Users struct {
	ID        string     `json:"id"`
	UserName  string     `json:"user_name"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}

//HashPassword this function is to bcrypt password
func (u *Users) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

//GenerateToken method to generate token
func (u *Users) GenerateToken() (*AuthToken, error) {

	expiredAt := time.Now().Add(time.Hour * 24 * 7) //one weeks token expired

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "graph",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil

}

//ComparePassword so you can loggin with your password whic mean decaode bcrypte
func (u *Users) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHasedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHasedPassword, bytePassword)
}
