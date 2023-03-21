package model

import "github.com/99designs/gqlgen/graphql"

type NewAnimeRankingInput struct {
	Title      string         `json:"title"`
	Rank       int            `json:"rank"`
	UserID     int            `json:"userID"`
	AnimeImage graphql.Upload `json:"animeImage"`
	// RelatedAnime []*string `json:"relatedAnime"`
}

type UserInformationInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
