package auth

import (
	"errors"

	"github.com/samber/lo"
)

const OWNER_SCOPE = "FRAYM_AUTH_OWNER"

var ErrUnauthorized = errors.New("unauthorized")

type Data struct {
	scopes   []string
	tenantId string
	userId   string
	data     map[string]any
}

func (d *Data) GetTenantId() string {
	return d.tenantId
}

func (d *Data) GetUserId() string {
	return d.userId
}

func (d *Data) GetScopes() []string {
	return d.scopes
}

func (d *Data) GetData() map[string]any {
	return d.data
}

func (d *Data) GetDataValue(key string) any {
	value, ok := d.data[key]
	if !ok {
		return nil
	}
	return value
}

func (d *Data) CheckAuthorization(permissions []string) error {
	if lo.Contains(d.scopes, OWNER_SCOPE) {
		return nil
	}

	if len(permissions) == 0 || !lo.Some(d.scopes, permissions) {
		return ErrUnauthorized
	}
	return nil
}
