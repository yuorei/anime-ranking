package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"context"
	"fmt"

	"github.com/yuorei/anime-ranking/application"
	"github.com/yuorei/anime-ranking/database/mysql"
	"github.com/yuorei/anime-ranking/database/table"
	"github.com/yuorei/anime-ranking/graph/model"
	"github.com/yuorei/anime-ranking/middlewares"
)

// CreateAnimeRanking is the resolver for the createAnimeRanking field.
func (r *mutationResolver) CreateAnimeRanking(ctx context.Context, input model.NewAnimeRankingInput) (*model.AnimeRanking, error) {
	customClaim := middlewares.CtxValue(ctx)

	result, err := application.AWSS3Upload(input.AnimeImage.File, input.AnimeImage.Filename)
	if err != nil {
		return nil, err
	}

	anime := table.AnimeRanking{
		UserID:        customClaim.ID,
		Title:         input.Title,
		Rank:          input.Rank,
		Description:   input.Description,
		AnimeImageURL: result.Location,
	}

	anime, err = mysql.InsertAnimeRanking(anime)
	if err != nil {
		return nil, err
	}

	return &model.AnimeRanking{
		AnimeID:       int(anime.ID),
		Title:         anime.Title,
		Rank:          anime.Rank,
		Description:   &anime.Description,
		AnimeImageURL: result.Location,
	}, nil
}

// UpdateAnimeRanking is the resolver for the updateAnimeRanking field.
func (r *mutationResolver) UpdateAnimeRanking(ctx context.Context, animeID int, input model.UpdateAnimeRankingInput) (*model.AnimeRanking, error) {
	customClaim := middlewares.CtxValue(ctx)
	oldAnime, err := mysql.GetAnimeRankingByID(animeID)
	if err != nil {
		return nil, err
	}

	if oldAnime.UserID != customClaim.ID {
		return nil, fmt.Errorf("userIDが違います")
	}

	if input.Title != nil {
		oldAnime.Title = *input.Title
	}
	if input.Rank != nil {
		oldAnime.Rank = *input.Rank
	}
	if input.Description != nil {
		oldAnime.Description = *input.Description
	}

	if input.AnimeImage != nil {
		result, err := application.AWSS3Upload(input.AnimeImage.File, input.AnimeImage.Filename)
		if err != nil {
			return nil, err
		}
		oldAnime.AnimeImageURL = result.Location
	}

	newAnime, err := mysql.UpdateAnimeRanking(animeID, oldAnime)
	if err != nil {
		return nil, err
	}

	return &model.AnimeRanking{
		AnimeID:       int(newAnime.ID),
		Title:         newAnime.Title,
		Description:   &newAnime.Description,
		AnimeImageURL: newAnime.AnimeImageURL,
		Rank:          newAnime.Rank,
		User:          nil,
	}, nil
}

// DeleteAnimeRanking is the resolver for the deleteAnimeRanking field.
func (r *mutationResolver) DeleteAnimeRanking(ctx context.Context, animeID int) (*model.DeletePayload, error) {
	customClaim := middlewares.CtxValue(ctx)
	anime, err := mysql.GetAnimeRankingByID(animeID)
	if err != nil {
		return &model.DeletePayload{
			Success: false,
		}, err
	}

	if anime.UserID != customClaim.ID {
		return &model.DeletePayload{
			Success: false,
		}, fmt.Errorf("userIDが違います")
	}
	delete, err := mysql.DeleteAnimeRanking(anime)
	if err != nil || delete {
		return &model.DeletePayload{
			Success: false,
		}, err
	}
	return &model.DeletePayload{
		Success: true,
	}, nil
}

// GetAnimeRanking is the resolver for the getAnimeRanking field.
func (r *queryResolver) GetAnimeRanking(ctx context.Context, id int) (*model.AnimeRanking, error) {
	anime, err := mysql.GetAnimeRankingByID(id)
	if err != nil {
		return nil, err
	}
	result := &model.AnimeRanking{
		AnimeID:       int(anime.ID),
		Title:         anime.Title,
		Rank:          anime.Rank,
		Description:   &anime.Description,
		AnimeImageURL: anime.AnimeImageURL,
	}
	return result, nil
}