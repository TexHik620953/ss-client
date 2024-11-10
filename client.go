package shadowsocks

import (
	"context"
	"fmt"
	"net"

	"github.com/TexHik620953/ss-client/core"
	"github.com/TexHik620953/ss-client/protocol/socks"
)

type Client struct {
	cipher core.Cipher
	addr   string
}

func New(addr, cipher, password string) (*Client, error) {

	ciph, err := core.PickCipher(cipher, []byte{}, password)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %v", cipher)
	}
	h := &Client{
		cipher: ciph,
		addr:   addr,
	}
	return h, nil
}

func (h *Client) Dialer(ctx context.Context, network, host string) (net.Conn, error) {
	addr := socks.ParseAddr(host)
	conn, err := net.Dial(network, h.addr)
	if err != nil {
		return nil, err
	}
	conn = h.cipher.StreamConn(conn)
	err = socks.WriteAddr(conn, addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
