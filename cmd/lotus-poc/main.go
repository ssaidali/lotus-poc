package main

import (
	"context"
	"log"
	"github.com/ssaidali/lotus-poc/lotus"
)

func main() {
	api, err := lotus.NewLotusRPC(context.Background(), "https://node.glif.io/space07/lotus/rpc/v0", http.Header{})
	if err != nil {
		log.Error(err)
	}
	defer api.Close()
}
