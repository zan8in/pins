package iputil

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	// ipv4c := "60.10.116.10"
	// ipv4a := "192.168.88.110"
	// ipv4b := "60.10.116.10/24"
	domain := "hackerone.com"
	// domain2 := "www.lf.gov.cn"
	// domain3 := "qq.com"

	// for _, ip := range ipv4 {
	// 	fmt.Println(ip)
	// }
	// netip, rangeNet, err := net.ParseCIDR(ipv4b)
	// fmt.Println(netip, rangeNet, err)

	fmt.Println(ToFQDN(domain))
}
