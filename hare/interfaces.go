package hare

import (
	"context"

	"github.com/spacemeshos/go-spacemesh/common/types"
)

//go:generate mockgen -package=mocks -destination=./mocks/mocks.go -source=./interfaces.go

type layerPatrol interface {
	SetHareInCharge(types.LayerID)
}

// Rolacle is the roles oracle provider.
type Rolacle interface {
	Validate(ctx context.Context, layer types.LayerID, round uint32, committeeSize int, id types.NodeID, sig []byte, eligibilityCount uint16) (bool, error)
	CalcEligibility(ctx context.Context, layer types.LayerID, round uint32, committeeSize int, id types.NodeID, sig []byte) (uint16, error)
	Proof(ctx context.Context, layer types.LayerID, round uint32) ([]byte, error)
	IsIdentityActiveOnConsensusView(ctx context.Context, edID string, layer types.LayerID) (bool, error)
}

type meshProvider interface {
	// LayerBlocks returns the blocks in a layer
	LayerBlocks(types.LayerID) ([]*types.Block, error)
	// GetBlock returns the block with the specified block ID
	GetBlock(types.BlockID) (*types.Block, error)
	// HandleValidatedLayer receives Hare output when it succeeds
	HandleValidatedLayer(ctx context.Context, validatedLayer types.LayerID, layer []types.BlockID)
	// InvalidateLayer receives the signal that Hare failed for a layer
	InvalidateLayer(ctx context.Context, layerID types.LayerID)
	// RecordCoinflip records the weak coinflip result for a layer
	RecordCoinflip(ctx context.Context, layerID types.LayerID, coinflip bool)
}