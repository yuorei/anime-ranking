package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"

	"github.com/yuorei/anime-ranking/auth"
	"github.com/yuorei/anime-ranking/graph/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	user, err := auth.GetUserByName(input.Name)
	if err != nil {
		return &model.LoginResponse{Success: false}, nil
	}

	if !auth.VerifyPassword(user.Password, input.Password) {
		return &model.LoginResponse{Success: false}, nil
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return &model.LoginResponse{Success: false}, nil
	}

	return &model.LoginResponse{
		Success: true,
		Token:   &token,
	}, nil
}
