package main

import (
	"context"
	"fmt"
	"net"
	"net/netip"
)

func main() {
	var oldIP net.IP
	var newIP netip.Addr
	var resolver net.Resolver

	// Bytes to IPv4
	oldIP = net.IPv4(10, 20, 30, 46)
	newIP = netip.AddrFrom4([4]byte{10, 30, 40, 50})
	fmt.Println("oldIP: ", oldIP)
	fmt.Println("newIP: ", newIP)

	// Host/IP resolution only in net.IP.
	host := "www.golang.org"
	resolvedHost, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("error resolving host")
	}
	fmt.Printf("resolved host %s to IPs %s\n", host, resolvedHost)

	resolvedIP, err := net.LookupHost(resolvedHost[0].String())
	if err != nil {
		fmt.Println("error resolving IP")
	}
	fmt.Printf("resolved IP %s to host %s\n", resolvedIP[0], host)

	// New Name Resolution in net package.
	ctx := context.Background()
	ipv4Resolved, err := resolver.LookupNetIP(ctx, "ip4", host)
	if err != nil {
		fmt.Printf("error with LookupNetIP: %s\n", err)
	}
	ipv6Resolved, err := resolver.LookupNetIP(ctx, "ip6", host)
	if err != nil {
		fmt.Printf("error with LookupNetIP: %s\n", err)
	}
	fmt.Println("ipv4Resolved: ", ipv4Resolved)
	//_, hackedIPv4, _ := strings.Cut(ipv4Resolved[0].String(), "f:")
	//fmt.Println("hacked IPv4: ", hackedIPv4)
	fmt.Println("ipv6Resolved: ", ipv6Resolved)
}
