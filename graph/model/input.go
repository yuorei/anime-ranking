package model

type NewAnimeRankingInput struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
	Rank     int    `json:"rank"`
	UserID   string `json:"userID"`
}

type UserInformationInput struct {
	Name    string `json:"name"`
	Password string `json:"password"`
}