package upgrader

import (
	"fmt"
	"github.com/mikioh/tcpinfo"

	"github.com/libp2p/go-libp2p-core/mux"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/transport"
)

type transportConn struct {
	mux.MuxedConn
	network.ConnMultiaddrs
	network.ConnSecurity
	transport transport.Transport
	scope     network.ConnManagementScope
	stat      network.ConnStats
	tconn     transport.TracingConn
}

var _ transport.CapableConn = &transportConn{}

func (t *transportConn) Transport() transport.Transport {
	return t.transport
}

func (t *transportConn) String() string {
	ts := ""
	if s, ok := t.transport.(fmt.Stringer); ok {
		ts = "[" + s.String() + "]"
	}
	return fmt.Sprintf(
		"<stream.Conn%s %s (%s) <-> %s (%s)>",
		ts,
		t.LocalMultiaddr(),
		t.LocalPeer(),
		t.RemoteMultiaddr(),
		t.RemotePeer(),
	)
}

func (t *transportConn) Stat() network.ConnStats {
	return t.stat
}

func (t *transportConn) Scope() network.ConnScope {
	return t.scope
}

func (t *transportConn) Close() error {
	defer t.scope.Done()
	return t.MuxedConn.Close()
}

func (t *transportConn) GetTCPInfo() (*tcpinfo.Info, error) {
	return t.tconn.GetTCPInfo()
}
