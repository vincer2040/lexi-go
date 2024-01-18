package lexigo

import (
	"io"
	"net"
	"time"

	"github.com/vincer2040/lexi-go/internal/builder"
	"github.com/vincer2040/lexi-go/internal/parser"
	lexidata "github.com/vincer2040/lexi-go/pkg/lexi-data"
)

type Client struct {
	addr        string
	connection  net.Conn
	isconnected bool
	builder     builder.Builder
}

func New(addr string) Client {
	return Client{addr, nil, false, builder.New()}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.connection = conn
	c.isconnected = true
	return nil
}

func (c *Client) IsConnected() bool {
	return c.isconnected
}

func (c *Client) Ping() (*lexidata.LexiData, error) {
	buf := c.builder.Reset().AddPing().Out()
	return c.send(buf)
}

func (c *Client) Auth(username string, password string) (*lexidata.LexiData, error) {
	buf := c.builder.Reset().AddArray(3).AddString("AUTH").AddString(username).AddString(password).Out()
	return c.send(buf)
}

func (c *Client) Set(key string, value string) (*lexidata.LexiData, error) {
	buf := c.builder.Reset().AddArray(3).AddString("SET").AddString(key).AddString(value).Out()
	return c.send(buf)
}

func (c *Client) Get(key string) (*lexidata.LexiData, error) {
	buf := c.builder.Reset().AddArray(2).AddString("GET").AddString(key).Out()
	return c.send(buf)
}

func (c *Client) Del(key string) (*lexidata.LexiData, error) {
	buf := c.builder.Reset().AddArray(2).AddString("DEL").AddString(key).Out()
	return c.send(buf)
}

func (c *Client) Close() {
	c.connection.Close()
}

func (c *Client) readFromSocket() ([]byte, int, error) {
	var res []byte
	amt_read := 0
	for {
		buf := make([]byte, 1024)
		c.connection.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
		n, err := c.connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else if err.(*net.OpError).Timeout() {
				if amt_read > 0 {
					break
				}
				continue
			}
			return nil, 0, err
		}
		res = append(res, buf...)
		amt_read += n
	}
	return res, amt_read, nil
}

func (c *Client) send(buf []byte) (*lexidata.LexiData, error) {
	_, err := c.connection.Write(buf)
	if err != nil {
		return nil, err
	}
	read, amt_read, err := c.readFromSocket()
	if err != nil {
		return nil, err
	}
	p := parser.New(read, amt_read)
	return p.Parse()
}
