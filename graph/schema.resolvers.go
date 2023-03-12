package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"

	"github.com/yuorei/anime-ranking/graph/model"
)

// RelatedAnime is the resolver for the relatedAnime field.
func (r *animeInformationResolver) RelatedAnime(ctx context.Context, obj *model.AnimeInformation) ([]*model.AnimeInformation, error) {
	panic(fmt.Errorf("not implemented: RelatedAnime - relatedAnime"))
}

// RegisterUser is the resolver for the registerUser field.
func (r *animeInformationResolver) RegisterUser(ctx context.Context, obj *model.AnimeInformation) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: RegisterUser - registerUser"))
}

// AnimeInformation is the resolver for the animeInformation field.
func (r *animeRankingResolver) AnimeInformation(ctx context.Context, obj *model.AnimeRanking) (*model.AnimeInformation, error) {
	panic(fmt.Errorf("not implemented: AnimeInformation - animeInformation"))
}

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.UserInformationInput) (*model.UserPayload, error) {
	url := "urlだよ"
	userPayload := &model.UserPayload{
		Name:           input.Naem,
		Password:       input.Password,
		ProfieImageURL: &url,
	}
	user := &model.User{
		Name:           userPayload.Name,
		Password:       userPayload.Password,
		UserID:         "ランダム",
		ProfieImageURL: userPayload.ProfieImageURL,
	}
	r.Resolver.users = append(r.Resolver.users, user)
	return userPayload, nil
}

// RegisterUserAnimeRanking is the resolver for the registerUserAnimeRanking field.
func (r *mutationResolver) RegisterUserAnimeRanking(ctx context.Context, input model.NewAnimeRankingInput) (*model.AnimeRankingPayload, error) {
	panic(fmt.Errorf("not implemented: RegisterUserAnimeRanking - registerUserAnimeRanking"))
}

// GetUserInformation is the resolver for the GetUserInformation field.
func (r *queryResolver) GetUserInformation(ctx context.Context) ([]*model.User, error) {
	// panic(fmt.Errorf("not implemented: GetUserInformation - GetUserInformation"))
	return r.users, nil
}

// GetAnimeRanking is the resolver for the GetAnimeRanking field.
func (r *queryResolver) GetAnimeRanking(ctx context.Context) ([]*model.AnimeRanking, error) {
	panic(fmt.Errorf("not implemented: GetAnimeRanking - GetAnimeRanking"))
	// return r.db, nil
}

// HaveAnime is the resolver for the haveAnime field.
func (r *userResolver) HaveAnime(ctx context.Context, obj *model.User) ([]*model.AnimeRanking, error) {
	panic(fmt.Errorf("not implemented: HaveAnime - haveAnime"))
}

// RelatedAnime is the resolver for the relatedAnime field.
func (r *newAnimeRankingInputResolver) RelatedAnime(ctx context.Context, obj *model.NewAnimeRankingInput, data []*string) error {
	panic(fmt.Errorf("not implemented: RelatedAnime - relatedAnime"))
}

// AnimeInformation returns AnimeInformationResolver implementation.
func (r *Resolver) AnimeInformation() AnimeInformationResolver { return &animeInformationResolver{r} }

// AnimeRanking returns AnimeRankingResolver implementation.
func (r *Resolver) AnimeRanking() AnimeRankingResolver { return &animeRankingResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// NewAnimeRankingInput returns NewAnimeRankingInputResolver implementation.
func (r *Resolver) NewAnimeRankingInput() NewAnimeRankingInputResolver {
	return &newAnimeRankingInputResolver{r}
}

type animeInformationResolver struct{ *Resolver }
type animeRankingResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type newAnimeRankingInputResolver struct{ *Resolver }
