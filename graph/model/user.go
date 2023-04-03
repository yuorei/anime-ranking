package model

type User struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Password       string          `json:"password"`
	ProfieImageURL string          `json:"profieImageURL"`
	Description    *string         `json:"description"`
	HaveAnime      []*AnimeRanking `json:"haveAnime"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }
