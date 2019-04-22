package utils

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	core "github.com/libp2p/go-libp2p-core/transport/test"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.Subtests instead.
var Subtests = core.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestTransport instead.
func SubtestTransport(t *testing.T, ta, tb transport.Transport, addr string, peerA peer.ID) {
	core.SubtestTransport(t, ta, tb, addr, peerA)
}
