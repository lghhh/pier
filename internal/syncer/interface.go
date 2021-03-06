package syncer

import "github.com/meshplus/bitxhub-model/pb"

type IBTPHandler func(ibtp *pb.IBTP)

//go:generate mockgen -destination mock_syncer/mock_syncer.go -package mock_syncer -source interface.go
type Syncer interface {
	// Start starts the service of syncer
	Start() error

	// Stop stops the service of syncer
	Stop() error

	RegisterIBTPHandler(handler IBTPHandler) error
}
