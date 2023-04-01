package model

import "github.com/99designs/gqlgen/graphql"

type UserInformationInput struct {
	Name        string         `json:"name"`
	Password    string         `json:"password"`
	Description string         `json:"description"`
	ProfieImage graphql.Upload `json:"profieImage"`
}

type UpdateUserInput struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	ProfieImage *graphql.Upload `json:"profieImage"`
}
