package model

type LoginInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool    `json:"success"`
	Token   *string `json:"token"`
}

type AuthOps struct {
	Login *LoginResponse `json:"login"`
}
