package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/sethj12/gqlgen-todos/graph/generated"
	"github.com/sethj12/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
	}
	r.todo = append(r.todo, todo)
	return todo, nil
}

func (r *mutationResolver) CreateNewUsers(ctx context.Context, input *model.RegisterInput) (*model.Register, error) {
	user := &model.Register{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Username: input.Username,
		Lastname: input.Lastname,
		Password: input.Password,
	}
	r.user = append(r.user, user)
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input *model.LoginInput) (string, error) {
	for _, v := range r.user {
		if v.Username == input.Username && v.Password == input.Password {
			return "successfully login", nil
		}
	}
	return "No user found", nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUserInput) (string, error) {
	for _, i := range r.user {
		fmt.Println(i)

	}

	return "blankkkkkkk", nil
}

func (r *queryResolver) Todo(ctx context.Context) ([]*model.Todo, error) {
	return r.todo, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.Register, error) {
	return r.user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateUserss(ctx context.Context, input *model.RegisterInput) (*model.Register, error) {
	user := &model.Register{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Lastname: input.Lastname,
		Password: input.Password,
	}
	r.user = append(r.user, user)
	return user, nil
}
func (r *mutationResolver) CreateNewUserss(ctx context.Context, input *model.RegisterInput) (*model.Register, error) {
	user := &model.Register{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Lastname: input.Lastname,
		Password: input.Password,
	}
	r.user = append(r.user, user)
	return user, nil
}
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.RegisterInput) (*model.Register, error) {
	panic(fmt.Errorf("not implemented"))
}
