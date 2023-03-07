package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"

	"github.com/yuorei/anime-ranking/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.UserInformationInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: RegisterUser - registerUser"))
}

// CreateAnimeRanking is the resolver for the createAnimeRanking field.
func (r *mutationResolver) CreateAnimeRanking(ctx context.Context, input model.NewAnimeRankingInput) (*model.AnimeRanking, error) {
	rankings := &model.AnimeRanking{
		AnimeInformation: &model.AnimeInformation{AnimeID: "hoge", Title: input.Title, ImageURL: input.ImageURL},
		Rank:             input.Rank,
		User:             &model.User{UserID: input.UserID},
	}
	r.db = append(r.db, rankings)
	return rankings, nil
}

// GetUserInformation is the resolver for the GetUserInformation field.
func (r *queryResolver) GetUserInformation(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUserInformation - GetUserInformation"))
}

// GetAnimeRanking is the resolver for the GetAnimeRanking field.
func (r *queryResolver) GetAnimeRanking(ctx context.Context) ([]*model.AnimeRanking, error) {
	// panic(fmt.Errorf("not implemented: GetAnimeRanking - GetAnimeRanking"))
	return r.db, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
