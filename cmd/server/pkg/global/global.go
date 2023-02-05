// package is not thread safe must be used to read service not write once initialized
package global

import "github.com/harshrastogiexe/cmd/server/pkg/core"

type Container[K any] map[core.ServiceToken]K

var container Container[any]

func init() {
	container = Container[any]{}
}

// add value to global registry
func Add[K any](token core.ServiceToken, item K) {
	container[token] = item
}

// gets value from global registry
func Get[T any](token core.ServiceToken) (T, bool) {
	value, found := container[token]
	return value.(T), found
}
