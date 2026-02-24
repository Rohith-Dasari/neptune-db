package main

import (
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

func main() {
	// Creating the connection to the server.
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection("wss://{{endpoint_name}}/gremlin",
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Cleanup
	defer driverRemoteConnection.Close()

	// Creating graph traversal
	g := gremlingo.Traversal_().With(driverRemoteConnection)

	// Perform traversal
	results, err := g.V().Limit(2).ToList()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print results
	for _, r := range results {
		fmt.Println(r.GetString())
	}
}
