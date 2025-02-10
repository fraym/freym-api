package auth

import (
	"fmt"
	"time"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

var (
	TokenTypeMisc    = TokenType("token")
	TokenTypeAccess  = TokenType("access_token")
	TokenTypeRefresh = TokenType("refresh_token")
)

type FraymClaims struct {
	jwt.RegisteredClaims
	Type     TokenType      `json:"type,omitempty"`
	TenantID string         `json:"tenantId,omitempty"`
	Scopes   []string       `json:"scopes,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
}

func ValidateTokenAndGetData(token string) (*Data, error) {
	authSecret := env.String("AUTH_SECRET")

	jwtToken, err := jwt.ParseWithClaims(token, &FraymClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(authSecret), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*FraymClaims)
	if !ok {
		return nil, fmt.Errorf("cannot parse jwt data, got claims of type %T", jwtToken.Claims)
	}

	return &Data{
		scopes:   claims.Scopes,
		tenantId: claims.TenantID,
		userId:   claims.Subject,
		data:     claims.Data,
	}, nil
}
