package main

import (
	"fmt"

	"github.com/zan8in/pins/netx"
)

func main() {
	conf := netx.Config{}
	tcpclient, err := netx.NewClient("139.224.106.219:2181", conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer tcpclient.Close()

	err = tcpclient.TcpSend([]byte("envi\r\nquit\r\n"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := tcpclient.TcpRecv()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))
	// testCases := []string{"3306", "21", "80", "8080"}
	// testCases := []string{"3306", "21", "80", "8080", "135", "137", "138", "445"}
	// for _, tc := range testCases {
	// 	conf := netutil.Config{
	// 		Network: "udp",
	// 	}

	// 	client, err := netutil.NewClient("127.0.0.1:"+tc, conf)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		continue
	// 	}
	// 	defer client.Close()

	// 	if err = client.Send([]byte("\n")); err != nil {
	// 		fmt.Println(err.Error())
	// 		continue
	// 	}

	// 	data, err := client.Receive()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		continue
	// 	}

	// 	fmt.Println(string(data) + "\n=======================\n")
	// 	time.Sleep(time.Second)
	// }
}
