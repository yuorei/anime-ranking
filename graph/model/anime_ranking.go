package model

type AnimeRanking struct {
	AnimeID       int     `json:"animeID"`
	Title         string  `json:"title"`
	Description   *string `json:"description"`
	AnimeImageURL string  `json:"animeImageURL"`
	Rank          int     `json:"rank"`
	User          *User   `json:"user"`
}
