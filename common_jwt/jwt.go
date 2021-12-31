package common_jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type User struct {
	Username          string `json:"username"`
	UserId            int    `json:"id"`
	ParentId          int    `json:"parent_id"`
	DefaultBusinessId int    `json:"default_business_id"`
}
type Claims struct {
	User
	jwt.StandardClaims
}

func JWTCreateToken(user User, securitySalt string) (tokenString string, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	// Create the JWT key used to create the signature
	mySigningKey := []byte(securitySalt)

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(60 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err = token.SignedString(mySigningKey)
	if err != nil {
		log.Errorf("Failed to create jwk token: %s", err.Error())
	}

	return
}

func JWTValidateToken(tokenString string, securitySalt string) (user User, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(securitySalt), nil
	})

	if claims, ok := token.Claims.(*Claims); !ok || !token.Valid {
		log.Errorf("Failed to validate jwk token: %s", err.Error())
		return
	} else {
		user = claims.User
	}
	return
}

func JWTRefreshToken(tokenString string, securitySalt string) (newTokenString string, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(securitySalt), nil
	})

	var user User
	if claims, ok := token.Claims.(*Claims); !ok || !token.Valid {
		// token is not valid. It may have expired
		// refresh only works on valid token
		log.Errorf("Failed to validate jwk token: %s", err.Error())
		return
	} else {
		user = claims.User
	}

	// create new token
	newTokenString, err = JWTCreateToken(user, securitySalt)
	if err != nil {
		log.Errorf("Failed to create new token during refresh jwk token: %s", err.Error())
		return
	}
	return
}
