package model

type Node interface {
	IsNode()
	// 任意のID
	GetID() string
}
