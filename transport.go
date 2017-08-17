package transport

import (
	"context"
	"io"
	"net"
	"time"

	logging "github.com/ipfs/go-log"
	smux "github.com/libp2p/go-stream-muxer"
	ma "github.com/multiformats/go-multiaddr"
)

var log = logging.Logger("transport")

// Conn is an extension of the net.Conn interface that provides multiaddr
// information, and an accessor for the transport used to create the conn
type Conn interface {
	io.Closer

	RemoteAddr() net.Addr
	RemoteMultiaddr() ma.Multiaddr

	LocalAddr() net.Addr
	LocalMultiaddr() ma.Multiaddr

	Transport() Transport
}

// A SingleStreamConn is a connection that provides a single channel between two endpoints
// e.g. a TCP connection
type SingleStreamConn interface {
	Conn

	io.Reader
	io.Writer

	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
}

// A MultiStreamConn is a connection that supports transport-level stream multiplexing.
// e.g. a QUIC connection
// The MultiStreamConn combines the smux.Conn and the Conn interface
// (unfortunately Go still doesn't allow duplicate interface methods...)
type MultiStreamConn interface {
	smux.Conn

	RemoteAddr() net.Addr
	RemoteMultiaddr() ma.Multiaddr

	LocalAddr() net.Addr
	LocalMultiaddr() ma.Multiaddr

	Transport() Transport
}

// Transport represents any device by which you can connect to and accept
// connections from other peers. The built-in transports provided are TCP and UTP
// but many more can be implemented, sctp, audio signals, sneakernet, UDT, a
// network of drones carrying usb flash drives, and so on.
type Transport interface {
	Dialer(laddr ma.Multiaddr, opts ...DialOpt) (Dialer, error)
	Listen(laddr ma.Multiaddr) (Listener, error)
	Matches(ma.Multiaddr) bool
}

// Dialer is an abstraction that is normally filled by an object containing
// information/options around how to perform the dial. An example would be
// setting TCP dial timeout for all dials made, or setting the local address
// that we dial out from.
type Dialer interface {
	Dial(raddr ma.Multiaddr) (Conn, error)
	DialContext(ctx context.Context, raddr ma.Multiaddr) (Conn, error)
	Matches(ma.Multiaddr) bool
}

// Listener is an interface closely resembling the net.Listener interface.  The
// only real difference is that Accept() returns Conn's of the type in this
// package, and also exposes a Multiaddr method as opposed to a regular Addr
// method
type Listener interface {
	Accept() (Conn, error)
	Close() error
	Addr() net.Addr
	Multiaddr() ma.Multiaddr
}

// DialOpt is an option used for configuring dialer behaviour
type DialOpt interface{}

type ReuseportOpt bool

var ReusePorts ReuseportOpt = true
