package response

import "github.com/RaymondSalim/API-server-template/server/models"

type FooResponse struct {
	Status string     `json:"status"`
	Foo    models.Foo `json:"foo,omitempty"`
}
