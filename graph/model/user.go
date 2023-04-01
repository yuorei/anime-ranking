package model

type User struct {
	UserID         int             `json:"userID"`
	Name           string          `json:"name"`
	Password       string          `json:"password"`
	ProfieImageURL string          `json:"profieImageURL"`
	Description    *string         `json:"description"`
	HaveAnime      []*AnimeRanking `json:"haveAnime"`
}
