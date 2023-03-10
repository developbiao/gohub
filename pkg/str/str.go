package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural convert word to plural user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular convert plural word to singular users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake convert snake_case, example: TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel covert CamelCase, example: _Topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel convert lowerCameCase, example TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
