package utils

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	"github.com/libp2p/go-libp2p-testing/suites/transport"

	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.SubtestProtocols instead.
func SubtestProtocols(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	ttransport.SubtestProtocols(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.SubtestBasic instead.
func SubtestBasic(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	ttransport.SubtestBasic(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.SubtestPingPong instead.
func SubtestPingPong(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	ttransport.SubtestPingPong(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-testing/suites/transport.SubtestCancel instead.
func SubtestCancel(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	ttransport.SubtestCancel(t, ta, tb, maddr, peerA)
}
