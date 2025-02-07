package crud

import (
	"github.com/Becklyn/go-wire-core/json"

	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
)

type AuthData struct {
	TenantId string
	Scopes   []string
	Data     map[string]any
}

func (a *AuthData) getProtobufAuthData() (*deliverypb.AuthData, error) {
	data := map[string]string{}

	for key, value := range a.Data {
		byteValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		data[key] = string(byteValue)
	}

	return deliverypb.AuthData_builder{
		TenantId: a.TenantId,
		Scopes:   a.Scopes,
		Data:     data,
	}.Build(), nil
}
