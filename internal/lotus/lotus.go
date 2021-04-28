package lotus

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
)

// LotusAPI tells the rpc client how to handle the different methods
type LotusAPI struct {
	Methods struct {
		ChainHead                         func(context.Context) (*TipSet, error)
		GasEstimateMessageGas             func(context.Context, *Message, *MessageSendSpec, TipSetKey) (*Message, error)
		StateGetActor                     func(context.Context, address.Address, TipSetKey) (*Actor, error)
		MpoolPush                         func(context.Context, *SignedMessage) (cid.Cid, error)
		StateWaitMsg                      func(context.Context, cid.Cid, uint64) (*MsgLookup, error)
		StateAccountKey                   func(context.Context, address.Address, TipSetKey) (address.Address, error)
		StateLookupID                     func(context.Context, address.Address, TipSetKey) (address.Address, error)
		StateReadState                    func(context.Context, address.Address, TipSetKey) (*ActorState, error)
		StateNetworkVersion               func(context.Context, TipSetKey) (network.Version, error)
		StateMarketBalance                func(context.Context, address.Address, TipSetKey) (MarketBalance, error)
		StateDealProviderCollateralBounds func(context.Context, abi.PaddedPieceSize, bool, TipSetKey) (DealCollateralBounds, error)
		StateMinerInfo                    func(context.Context, address.Address, TipSetKey) (MinerInfo, error)
		StateMinerProvingDeadline         func(context.Context, address.Address, TipSetKey) (*dline.Info, error)
		StateCall                         func(context.Context, *Message, TipSetKey) (*InvocResult, error)
		ChainReadObj                      func(context.Context, cid.Cid) ([]byte, error)
		ChainGetMessage                   func(context.Context, cid.Cid) (*Message, error)
	}
	closer jsonrpc.ClientCloser
}

// NewLotusRPC starts a new lotus RPC client
func NewLotusRPC(ctx context.Context, addr string, header http.Header) (API, error) {
	var res LotusAPI
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin",
		[]interface{}{
			&res.Methods,
		},
		header,
	)
	res.closer = closer
	return &res, err
}

func (a *LotusAPI) Close() {
	a.closer()
}

func (a *LotusAPI) ChainHead(ctx context.Context) (*TipSet, error) {
	return a.Methods.ChainHead(ctx)
}