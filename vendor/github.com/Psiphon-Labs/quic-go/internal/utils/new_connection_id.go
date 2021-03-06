package utils

import (
	"github.com/Psiphon-Labs/quic-go/internal/protocol"
)

// NewConnectionID is a new connection ID
type NewConnectionID struct {
	SequenceNumber      uint64
	ConnectionID        protocol.ConnectionID
	StatelessResetToken *[16]byte
}
