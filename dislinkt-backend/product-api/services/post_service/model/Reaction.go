package model

type Reaction struct {
	UserName string       `bson:"user_id"`
	Reaction ReactionType `bson:"reaction"`
}
