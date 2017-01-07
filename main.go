package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgraph/goclient/client"
	"github.com/dgraph-io/dgraph/query/graph"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("DialTCPConnection")
	}
	defer conn.Close()

	c := graph.NewDgraphClient(conn)

	// Running a mutation.
	q := "mutation { set {<alice> <follows> <bob> . \n <alice> <name> \"Alice\" . \n <bob> <name> \"Bob\" . }}"
	req := client.Req{}
	req.SetQuery(q)
	c.Run(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}

	// Running a query.
	q = `{me(_xid_: alice) { name _xid_ follows { name _xid_ }  }}`
	req := client.Req{}
	req.SetQuery(q)
	resp, err = c.Run(context.Background(), req.Request())
	if err != nil {
		log.Fatalf("Error in getting response from server, %s", err)
	}
	root := resp.N
	fmt.Printf("Subgraph %+v\n", root)
}
