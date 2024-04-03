package sequencesender

import (
	"context"

	ethmanTypes "github.com/0xPolygonHermez/zkevm-sequence-sender/etherman/types"
	"github.com/ethereum/go-ethereum/common"
)

type ethermaner interface {
	CurrentNonce(ctx context.Context, account common.Address) (uint64, error)
	BuildSequenceBatchesTxData(sender common.Address, sequences []ethmanTypes.Sequence, maxSequenceTimestamp uint64, lastSequencedBatchNumber uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (to *common.Address, data []byte, err error)
	GetLatestBatchNumber() (uint64, error)
}

type dataAbilitier interface {
	PostSequence(ctx context.Context, sequences []ethmanTypes.Sequence) ([]byte, error)
}
