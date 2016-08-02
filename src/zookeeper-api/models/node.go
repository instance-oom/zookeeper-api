package models

// ZNode - zookeeper node info
type ZNode struct {
	Path  string `json:'Path'`
	Value string `json:'value'`
}
