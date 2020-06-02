package tokenservice

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	user "github.com/paulcockrell/shippy/services/user/proto/user"
	repository "github.com/paulcockrell/shippy/services/user/repository"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

// TokenService -
type TokenService struct {
	Repository repository.Repository
}

// CustomClaims - our custom metadata, which will be hashed
// and sent as the second segment in the JWT
type CustomClaims struct {
	User *user.User
	jwt.StandardClaims
}

// Authable -
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *user.User) (string, error)
}

// Decode a token string into a token object
func (e *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (e *TokenService) Encode(user *user.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "com.foo.service.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}
