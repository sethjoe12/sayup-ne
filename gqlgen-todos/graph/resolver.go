package graph

import "github.com/sethj12/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todo []*model.Todo
	user []*model.Register
}
