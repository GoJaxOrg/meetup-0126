package main

import (
	"fmt"
	"net"
	"net/netip"
	//"net/netip"
)

/*
The contents are mutable. The underlying type of net.IP is []byte, which means that net.IP is passed by reference, and
any function that handles it can change its contents.
*/
func main() {
	oldAddr := net.IP{10, 10, 10, 0}
	fmt.Println(oldAddr.String())
	changeOld(oldAddr)

	//mask(oldAddr)
	//
	address, _ := netip.ParseAddr("10.20.30.40")
	fmt.Println(address.String())
	changeNew(address)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(address)
	testAddr := getIP()

	if testAddr == address {
		fmt.Println("addresses are equal")
	} else {
		fmt.Println("address mismatch")
	}

	//newOldIP := getOldIP()
	//if newOldIP == oldAddr {
	//
	//}
	var ipMap = make(map[netip.Addr]string)
	ipMap[address] = "http://mydomain.com"
	fmt.Println(ipMap)
}

func mask(addr net.IP) {
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
}

func changeOld(ip net.IP) {
	ip = net.IP{192, 168, 1, 18}
	fmt.Println(ip.String())
}

func changeNew(ip netip.Addr) {
	ip, _ = netip.ParseAddr("192.168.100.0")
	fmt.Println(ip.String())
}

func getOldIP() net.IP {
	return net.IPv4(1, 2, 3, 4)
}

func getIP() netip.Addr {
	return netip.MustParseAddr("1.2.3.4")
}
