package model

type NewAnimeRankingInput struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageURL"`
	Rank     int    `json:"rank"`
	UserID   string `json:"userID"`
}

type UserInformationInput struct {
	Naem     string `json:"naem"`
	Password string `json:"password"`
}