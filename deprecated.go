package transport

import moved "github.com/libp2p/go-libp2p-core/transport"

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.DialTimeout instead.
// Warning: it's not possible to alias variables in Go. Setting a value here may have no effect.
var DialTimeout = moved.DialTimeout

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.AcceptTimeout instead.
// Warning: it's not possible to alias variables in Go. Setting a value here may have no effect.
var AcceptTimeout = moved.AcceptTimeout

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.UpgradedConn instead.
type Conn = moved.UpgradedConn

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.Transport instead.
type Transport = moved.Transport

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.Listener instead.
type Listener = moved.Listener

// Deprecated: Use github.com/libp2p/go-libp2p-core/transport.TransportNetwork instead.
type Network = moved.TransportNetwork
