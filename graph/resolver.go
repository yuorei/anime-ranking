package graph

import "github.com/yuorei/anime-ranking/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	db []*model.AnimeRanking
	users []*model.User
}
