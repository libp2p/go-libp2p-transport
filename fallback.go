package transport

import (
	"fmt"

	utp "github.com/anacrolix/utp"
	ma "github.com/jbenet/go-multiaddr"
	manet "github.com/jbenet/go-multiaddr-net"
	mautp "github.com/jbenet/go-multiaddr-net/utp"
	mafmt "github.com/whyrusleeping/mafmt"
)

type FallbackDialer struct {
	madialer manet.Dialer
}

func (fbd *FallbackDialer) Matches(a ma.Multiaddr) bool {
	return mafmt.TCP.Matches(a) || mafmt.UTP.Matches(a)
}

func (fbd *FallbackDialer) Dial(a ma.Multiaddr) (Conn, error) {
	if mafmt.TCP.Matches(a) {
		return fbd.tcpDial(a)
	}
	if mafmt.UTP.Matches(a) {
		return fbd.utpDial(a)
	}
	return nil, fmt.Errorf("cannot dial %s with fallback dialer", a)
}

func (fbd *FallbackDialer) tcpDial(raddr ma.Multiaddr) (Conn, error) {
	var c manet.Conn
	var err error
	c, err = fbd.madialer.Dial(raddr)

	if err != nil {
		return nil, err
	}

	return &connWrap{
		Conn: c,
	}, nil
}

func (fbd *FallbackDialer) utpDial(raddr ma.Multiaddr) (Conn, error) {
	_, addr, err := manet.DialArgs(raddr)
	if err != nil {
		return nil, err
	}

	con, err := utp.Dial(addr)
	if err != nil {
		return nil, err
	}

	mnc, err := manet.WrapNetConn(&mautp.Conn{Conn: con})
	if err != nil {
		return nil, err
	}

	return &connWrap{
		Conn: mnc,
	}, nil
}
