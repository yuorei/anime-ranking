package model

type AnimeRanking struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Description   *string `json:"description"`
	AnimeImageURL string  `json:"animeImageURL"`
	Rank          int     `json:"rank"`
	User          *User   `json:"user"`
}

func (AnimeRanking) IsNode()            {}
func (this AnimeRanking) GetID() string { return this.ID }
