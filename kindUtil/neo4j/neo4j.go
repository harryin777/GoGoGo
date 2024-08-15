package neo4j

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Neo4J interface {
	init() neo4j.Driver
}

type Neo4jInstance struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewNeo4j() Neo4jInstance {
	return Neo4jInstance{
		Url:      "http://192.168.63.6:7474",
		Username: "",
		Password: "",
	}
}

func (n *Neo4jInstance) init() neo4j.Driver {
	driver, err := neo4j.NewDriver(n.Url, neo4j.BasicAuth(n.Username, n.Password, ""))
	if err != nil {
		panic(err)
	}
	return driver
}

func CloseDriver(driver neo4j.Driver) error {
	return driver.Close()
}
