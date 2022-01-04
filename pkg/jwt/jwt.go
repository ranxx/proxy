package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// DefaultMethod ...
var DefaultMethod = jwt.SigningMethodHS256

// Claim ...
type Claim struct {
	jwt.StandardClaims
	UserID   int64  `json:"userId,omitempty"`
	Acccount string `json:"account,omitempty"`
}

// Parse ...
func Parse(key []byte, tokenStr string, claim *Claim, method jwt.SigningMethod) error {
	if method == nil {
		method = DefaultMethod
	}

	token, err := jwt.ParseWithClaims(tokenStr, claim, func(token *jwt.Token) (interface{}, error) {
		if token.Method != method {
			return nil, jwt.ErrHashUnavailable
		}
		return key, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.NewValidationError("signature validation failed", jwt.ValidationErrorSignatureInvalid)
	}

	// 过期
	if !time.Unix(claim.ExpiresAt, 0).Local().After(time.Now()) {
		return jwt.NewValidationError("the token has expired", jwt.ValidationErrorExpired)
	}

	return nil
}

// Generate ...
func Generate(key []byte, claim *Claim, method jwt.SigningMethod) (string, error) {
	if method == nil {
		method = DefaultMethod
	}

	token := jwt.NewWithClaims(method, claim)

	return token.SignedString(key)
}
