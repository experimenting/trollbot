package troll

// VarRepository interface
type VarRepository interface {
	GetRandomUniqueVar(varType string) (string, error)
	GetAllTopics() []string
}

// FeedRepository interface
type FeedRepository interface {
	GetByContext(context string) (Feed, error)
	GetAllTags() []string
}
