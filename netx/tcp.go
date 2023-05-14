package netx

func NewTcpClient(address string, conf Config) (*Client, error) {
	conf.Network = "tcp"
	return NewClient(address, conf)
}

func (c *Client) TcpSend(data []byte) error {
	return c.Send(data)
}

func (c *Client) TcpRecv() ([]byte, error) {
	return c.Receive()
}
