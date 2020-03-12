package ethereum

import (
	"context"
	"errors"
	"math/big"

	"github.com/common"
	"github.com/core/types"
)

var NotFound = errors.New("not found")
type ChainReader interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)

	// This method subscribes to notifications about changes of the head block of
	// the canonical chain.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (Subscription, error)
}

// Subscription represents an event subscription where events are
// delivered on a data channel.
type Subscription interface {
	// Unsubscribe cancels the sending of events to the data channel
	// and closes the error channel.
	Unsubscribe()
	// Err returns the subscription error channel. The error channel receives
	// a value if there is an issue with the subscription (e.g. the network connection
	// delivering the events has been closed). Only one value will ever be sent.
	// The error channel is closed by Unsubscribe.
	Err() <-chan error
}
