package utils

import (
	"testing"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	moved "github.com/libp2p/go-libp2p-core/transport/test"

	ma "github.com/multiformats/go-multiaddr"
)

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.VerboseDebugging instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var VerboseDebugging = moved.VerboseDebugging

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.StressTestTimeout instead.
// Warning: it's impossible to alias a var in go, so reads and writes to this variable may be inaccurate
// or not have the intended effect.
var StressTestTimeout = moved.StressTestTimeout

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.Options instead.
type Options = moved.Options

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress instead.
func SubtestStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID, opt Options) {
	moved.SubtestStress(t, ta, tb, maddr, peerA, opt)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStreamOpenStress instead.
func SubtestStreamOpenStress(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStreamOpenStress(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStreamReset instead.
func SubtestStreamReset(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStreamReset(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress1Conn1Stream1Msg instead.
func SubtestStress1Conn1Stream1Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress1Conn1Stream1Msg(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress1Conn1Stream100Msg instead.
func SubtestStress1Conn1Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress1Conn1Stream100Msg(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress1Conn100Stream100Msg instead.
func SubtestStress1Conn100Stream100Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress1Conn100Stream100Msg(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress50Conn10Stream50Msg instead.
func SubtestStress50Conn10Stream50Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress50Conn10Stream50Msg(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress1Conn1000Stream10Msg instead.
func SubtestStress1Conn1000Stream10Msg(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress1Conn1000Stream10Msg(t, ta, tb, maddr, peerA)
}

// Deprecated: use github.com/libp2p/go-libp2p-core/transport/test.SubtestStress1Conn100Stream100Msg10MB instead.
func SubtestStress1Conn100Stream100Msg10MB(t *testing.T, ta, tb transport.Transport, maddr ma.Multiaddr, peerA peer.ID) {
	moved.SubtestStress1Conn100Stream100Msg10MB(t, ta, tb, maddr, peerA)
}
