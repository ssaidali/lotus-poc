package lotus

import (
	"context"
	"net/http"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/chain/types"
)

// LotusAPI tells the rpc client how to handle the different methods
type LotusAPI struct {
	node api.FullNode
	closer jsonrpc.ClientCloser
}

// NewLotusRPC starts a new lotus RPC client
func NewLotusRPC(ctx context.Context, addr string, header http.Header) (*LotusAPI, error) {
	var a LotusAPI
	node, closer, err := client.NewFullNodeRPCV1(ctx, addr, header)
	a.node = node
	a.closer = closer
	return &a, err
}

func (a *LotusAPI) Close() {
	a.closer()
}

func (a *LotusAPI) ChainHead(ctx context.Context) (*types.TipSet, error) {
	return a.node.ChainHead(ctx)
}