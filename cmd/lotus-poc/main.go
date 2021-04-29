package main

import (
	"context"
	"fmt"
	"net/http"
	"github.com/ssaidali/lotus-poc/internal/lotus"
)

func main() {
	api, err := lotus.NewLotusRPC(context.Background(), "https://node.glif.io/space07/lotus/rpc/v0", http.Header{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer api.Close()

	head, err := api.ChainHead(context.Background())
	if err != nil {
		fmt.Printf("Fatal %s", err)
	}
	fmt.Println(head.Height())
}
