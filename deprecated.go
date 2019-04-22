package transport

import core "github.com/libp2p/go-libp2p-core/transport"

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.DialTimeout instead.
// Warning: it's not possible to alias variables in Go. Setting a value here may have no effect.
var DialTimeout = core.DialTimeout

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.AcceptTimeout instead.
// Warning: it's not possible to alias variables in Go. Setting a value here may have no effect.
var AcceptTimeout = core.AcceptTimeout

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.UpgradedConn instead.
type Conn = core.UpgradedConn

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.Transport instead.
type Transport = core.Transport

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.Listener instead.
type Listener = core.Listener

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.TransportNetwork instead.
type Network = core.TransportNetwork
