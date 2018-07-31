package mapper

import (
	"reflect"
	"github.com/pkg/errors"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

//getNode return val as graph.Node if cast successful, otherwise return a detailed error.
func getNode(val interface{}) (graph.Node, error) {
	node, ok := val.(graph.Node)
	if !ok {
		var graphNode graph.Node
		return graphNode, castingError(graphNode, val)
	}
	return node, nil
}

//getNode return val as graph.Relationship if cast successful, otherwise return a detailed error.
func getRelationship(val interface{}) (graph.Relationship, error) {
	rel, ok := val.(graph.Relationship)
	if !ok {
		var graphRel graph.Relationship
		return graphRel, castingError(graphRel, rel)
	}
	return rel, nil
}

//getStringProperty return requested key value from map as a string. If key not found returns empty string and nil,
// returns casting error if val cannot be cast to string.
func getStringProperty(key string, props map[string]interface{}) (string, error) {
	val, ok := props[key]
	if !ok {
		return "", nil
	}

	strVal, ok := val.(string)
	if !ok {
		return "", castingError("", val)
	}
	return strVal, nil
}

//getint64Property return requested key value from map as a int64. If key not found returns empty 0 and nil,
// returns casting error if val cannot be cast to int64.
func getint64Property(key string, props map[string]interface{}) (int64, error) {
	val, ok := props[key]
	if !ok {
		return 0, nil
	}

	intVal, ok := val.(int64)
	if !ok {
		var expected int64
		return expected, castingError(expected, val)
	}
	return intVal, nil
}

func castingError(expected interface{}, actual interface{}) error {
	return errors.Errorf("failed to cast value to requested type, expected %q but was type %q", reflect.TypeOf(expected).String(), reflect.TypeOf(actual).String())
}
