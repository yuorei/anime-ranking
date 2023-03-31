package model

import "github.com/99designs/gqlgen/graphql"

type NewAnimeRankingInput struct {
	Title       string         `json:"title"`
	Rank        int            `json:"rank"`
	Description string         `json:"description"`
	AnimeImage  graphql.Upload `json:"animeImage"`
}

type UserInformationInput struct {
	Name        string         `json:"name"`
	Password    string         `json:"password"`
	Description string         `json:"description"`
	ProfieImage graphql.Upload `json:"profieImage"`
}

type UpdateAnimeRankingInput struct {
	Title       *string         `json:"title"`
	Description *string         `json:"description"`
	Rank        *int            `json:"rank"`
	AnimeImage  *graphql.Upload `json:"animeImage"`
}
