package utils

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	moved "github.com/libp2p/go-libp2p-core/transport/test"

	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestProtocols instead.
func SubtestProtocols(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestProtocols(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestBasic instead.
func SubtestBasic(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestBasic(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestPingPong instead.
func SubtestPingPong(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestPingPong(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestCancel instead.
func SubtestCancel(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestCancel(t, ta, tb, maddr, peerA)
}
