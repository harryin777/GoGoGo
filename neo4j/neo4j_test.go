package neo4j

import "testing"

func Test_init(t *testing.T) {

	n := NewNeo4j()
	n.init()
	select {}
}
