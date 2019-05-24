package utils

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	"github.com/libp2p/go-libp2p-testing/suites/transport"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.Subtests instead.
var Subtests = ttransport.Subtests

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.SubtestTransport instead.
func SubtestTransport(t *testing.T, ta, tb transport.Transport, addr string, peerA peer.ID) {
	ttransport.SubtestTransport(t, ta, tb, addr, peerA)
}
