package netx

func NewUdpClient(address string, conf Config) (*Client, error) {
	conf.Network = "udp"
	return NewClient(address, conf)
}

func (c *Client) UdpSend(data []byte) error {
	return c.Send(data)
}

func (c *Client) UdpRecv() ([]byte, error) {
	return c.Receive()
}
