// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AnimeInformation struct {
	AnimeID       string              `json:"animeID"`
	Title         string              `json:"title"`
	AnimeImageURL string              `json:"animeImageURL"`
	RelatedAnime  []*AnimeInformation `json:"relatedAnime"`
	RegisterUser  []*User             `json:"registerUser"`
}

type AnimeRanking struct {
	AnimeInformation *AnimeInformation `json:"animeInformation"`
	Rank             int               `json:"rank"`
}

type AnimeRankingPayload struct {
	Title         string    `json:"title"`
	Rank          int       `json:"rank"`
	AnimeImageURL *string   `json:"animeImageURL"`
	RelatedAnime  []*string `json:"relatedAnime"`
}

type User struct {
	UserID         string          `json:"userID"`
	Name           string          `json:"name"`
	Password       string          `json:"password"`
	ProfieImageURL *string         `json:"profieImageURL"`
	HaveAnime      []*AnimeRanking `json:"haveAnime"`
}

type UserPayload struct {
	Name           string  `json:"name"`
	Password       string  `json:"password"`
	ProfieImageURL *string `json:"profieImageURL"`
}
