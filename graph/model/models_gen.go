// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AnimeInformation struct {
	AnimeID      string              `json:"animeID"`
	Title        string              `json:"title"`
	ImageURL     string              `json:"imageURL"`
	RelatedAnime []*AnimeInformation `json:"relatedAnime"`
	RegisterUser []*User             `json:"registerUser"`
}

type AnimeRanking struct {
	AnimeInformation *AnimeInformation `json:"animeInformation"`
	Rank             int               `json:"rank"`
	User             *User             `json:"user"`
}

type User struct {
	UserID       string          `json:"userID"`
	Name         string          `json:"name"`
	Password     string          `json:"password"`
	RelatedAnime []*AnimeRanking `json:"relatedAnime"`
}
