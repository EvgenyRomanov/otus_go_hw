package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here.
	return &SimpleTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.

type SimpleTelnetClient struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func (t *SimpleTelnetClient) Connect() (err error) {
	t.connection, err = net.DialTimeout("tcp", t.address, t.timeout)
	return err
}

func (t *SimpleTelnetClient) Close() error {
	return t.connection.Close()
}

func (t *SimpleTelnetClient) Send() error {
	_, err := io.Copy(t.connection, t.in)
	return err
}

func (t *SimpleTelnetClient) Receive() error {
	_, err := io.Copy(t.out, t.connection)
	return err
}
