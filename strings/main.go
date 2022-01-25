// This example demonstrates the combined changes of the new strings.cut functionality with the netip package.
package main

import (
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"time"
)

func main() {
	var ipList = []string{
		"10.45.78.90:3343",
		"127.0.0.1:4000",
		"123.45.67.89:3306",
		"78.54.12.63:2345",
		"250.149.166.81:443",
		"208.41.196.190:3000",
		"10.10.38.65",
		"212.72.170.28:443",
		"250.66.16.72:30123",
		"207.38.75.46:8080",
		"158.24.28.232:26257",
	}

	var (
		addr, port string
		found      bool
		addrPorts  []netip.AddrPort
		tmpAddr    netip.Addr
		tmpPort    uint16
	)

	for _, ip := range ipList {
		addr, port, found = strings.Cut(ip, ":")
		switch found {
		case false:
			// If the separator isn't found, presume port 80.
			tmpAddr, _ = netip.ParseAddr(addr)
			tmpPort = uint16(80)
			addrPorts = append(addrPorts, netip.AddrPortFrom(tmpAddr, tmpPort))
		case true:
			tmpAddr, _ = netip.ParseAddr(addr)
			parsedPort, err := strconv.ParseUint(port, 10, 16)
			if err != nil {
				fmt.Println(err)
			}
			tmpPort = uint16(parsedPort)
			addrPorts = append(addrPorts, netip.AddrPortFrom(tmpAddr, tmpPort))
		}
	}
	for _, addrPort := range addrPorts {
		fmt.Printf("Type: %T; Data: %v\n", addrPort, addrPort)
		rs, err := net.DialTimeout("tcp", addrPort.String(), time.Millisecond*50)

		if err != nil {
			fmt.Println("\terror: ", err)
			continue
		}
		fmt.Printf("\tSuccessfully connected to %s\n", addrPort.String())
		err = rs.Close()
	}

	//var addrPorts []netip.AddrPort
	//
	//for _, ip := range ipList {
	//	addrPort, err := netip.ParseAddrPort(ip)
	//	if err != nil {
	//		fmt.Printf("error parsing ip address `%s` : %s\n", ip, err)
	//		continue
	//	}
	//	addrPorts = append(addrPorts, addrPort)
	//}
	//
	//for _, addrPort := range addrPorts {
	//	fmt.Printf("Type: %T; Data: %v\n", addrPort, addrPort)
	//	rs, err := net.DialTimeout("tcp", addrPort.String(), time.Millisecond*50)
	//
	//	if err != nil {
	//		fmt.Println("\terror: ", err)
	//		continue
	//	}
	//	fmt.Printf("\tSuccessfully connected to %s\n", addrPort.String())
	//	err = rs.Close()
	//}
}
