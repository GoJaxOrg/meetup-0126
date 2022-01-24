// This example demonstrates the combined changes of the new strings.cut functionality with the netip package.
package main

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"
)

func main() {
	var ipList = []string{
		"10.45.78.90:3343",
		"123.45.67.89:3306",
		"78.54.12.63:2345",
		"250.149.166.81:443",
		"208.41.196.190:3000",
		"10.10.38.65",
		"212.72.170.28:443",
		"250.66.16.72:3000",
		"207.38.75.46:23567",
		"158.24.28.232:30123",
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
			parsedPort, _ := strconv.ParseUint(port, 16, 16)
			tmpPort = uint16(parsedPort)
			addrPorts = append(addrPorts, netip.AddrPortFrom(tmpAddr, tmpPort))
		}
	}

	for _, addrPort := range addrPorts {
		fmt.Printf("Type: %T; Data: %v\n", addrPort, addrPort)
	}
}
