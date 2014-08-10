package main

import (
	"github.com/gocql/gocql"
	"log"
	"fmt"
)

func main() {
	cluster := gocql.NewCluster("localhost")

	cluster.Keyspace = "config"

	session, _ := cluster.CreateSession()

	defer session.Close()

	var property_name, property_value string

	if err := session.Query("select property_name, property_value from config").Scan(&property_name, &property_value); err != nil {
		log.Fatal(err)
	}

	fmt.Println(property_name, property_value)
}



