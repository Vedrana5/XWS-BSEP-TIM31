package model

type PostStatus int

const (
	LIKED PostStatus = iota
	DISLIKED
	NEUTRAL
)
