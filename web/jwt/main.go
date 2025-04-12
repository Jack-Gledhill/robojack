package jwt

import (
	"time"

	"github.com/Jack-Gledhill/robojack/config"
	"github.com/Jack-Gledhill/robojack/web/oauth"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Validity is the duration that a new JWT should be valid for
var Validity = time.Hour * 24 * time.Duration(config.Web.JWT.Validity)

// Claims defines the expected format that every JWT's claims should follow
type Claims struct {
	jwt.RegisteredClaims
	User *oauth.User `json:"user"`
}

// New creates a new, signed JWT for the given user
// This will add all the necessary claims for you, including the expiry date and a UUID v4
func New(u *oauth.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		User: u,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://github.com/Jack-Gledhill/robojack",
			Subject:   u.ID,
			Audience:  []string{"https://github.com/Jack-Gledhill/robojack"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(Validity)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	})

	return token.SignedString([]byte(config.Web.JWT.SigningSecret))
}

// Validate takes a JWT and performs various checks on its validity.
// This function will return true if ALL the following conditions are met:
// - Token was signed with the configured SigningSecret
// - Token's claims can be properly unmarshalled into Claims
// - Token's ExpiresAt is in the future
// - Token's NotBefore is in the past
func Validate(signedToken string) (bool, *Claims, error) {
	// Parse the token and check the signature
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Web.JWT.SigningSecret), nil
	})
	if err != nil {
		return false, nil, err
	}

	// Unmarshal the claims and make sure they're formatted as expected
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return false, nil, nil
	}

	// Check if the token has expired
	if claims.ExpiresAt.Before(time.Now()) {
		return false, nil, nil
	}

	// Check if we're not able to process the token yet
	if claims.NotBefore.After(time.Now()) {
		return false, nil, nil
	}

	return true, claims, nil
}
