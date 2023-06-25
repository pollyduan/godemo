package main

import (
	"fmt"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

func main() {
	var ip net.IP = net.ParseIP("172.16.100.0").To4()
	var mask = net.CIDRMask(23, 32)
	network := net.IPNet{
		IP: ip,
		// Mask: net.IPMask{255, 255, 254, 0},
		Mask: mask,
	}
	// 1个字节作为子掩码，因此可以划分为2个子网
	// 获取第一个子网
	// https://pkg.go.dev/net#CIDRMask
	i, _ := cidr.Subnet(&network, 1, 0)
	PrintSubnet(i)

	ipv4Addr := net.ParseIP("192.0.2.1")
	ipv4Mask := net.CIDRMask(23, 32)
	fmt.Printf("ipv4Mask: %v\n", ipv4Mask)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))

}

func PrintSubnet(n *net.IPNet) {
	fmt.Println("---")
	fmt.Printf("IP: %v\n", n.IP)
	fmt.Printf("Mask: %v\n", n.Mask)
	fmt.Printf("Network: %v\n", n.Network())
	fmt.Printf("String: %v\n", n.String())
	fmt.Printf("AddressCount: %v\n", cidr.AddressCount(n))
	i1, i2 := cidr.AddressRange(n)
	fmt.Printf("From: %v, To: %v \n", i1, i2)
	fmt.Println("---")
}
