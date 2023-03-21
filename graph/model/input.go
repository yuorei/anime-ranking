package model

type NewAnimeRankingInput struct {
	Title        string    `json:"title"`
	Rank         int       `json:"rank"`
	UserID       int       `json:"userID"`
	// RelatedAnime []*string `json:"relatedAnime"`
}

type UserInformationInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
