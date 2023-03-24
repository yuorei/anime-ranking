package model

import "github.com/99designs/gqlgen/graphql"

type NewAnimeRankingInput struct {
	Title      string         `json:"title"`
	Rank       int            `json:"rank"`
	AnimeImage graphql.Upload `json:"animeImage"`
}

type UserInformationInput struct {
	Name        string         `json:"name"`
	Password    string         `json:"password"`
	ProfieImage graphql.Upload `json:"profieImage"`
}
