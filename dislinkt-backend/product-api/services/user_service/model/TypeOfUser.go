package model

type TypeOfUser int

const (
	ADMIN TypeOfUser = iota
	REGISTERED_USER
	UNREGISTERED_USER
)
