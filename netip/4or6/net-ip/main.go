// This example demonstrates creating a large amount of IP addresses using the older net.IP method.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run() error {
	ipList := loadIPList("mapped-or-not.txt")
	identifyMappedAddresses(ipList, true)

	return nil
}

func identifyMappedAddresses(ipList []string, print bool) {
	var (
		ipAddresses []net.IP
		tip         net.IP
	)
	for _, ip := range ipList {
		tip = net.ParseIP(ip)
		ipAddresses = append(ipAddresses, tip)
	}

	if print {
		for i, nip := range ipAddresses {
			if isIPv6(nip.String()) {
				fmt.Printf("%d. %s\n", i+1, nip.String()+" *")
			} else {
				fmt.Printf("%d. %s\n", i+1, nip.String())
			}
		}
		fmt.Printf("\n%s\n", "* Denotes IPv4-mapped or IPv6 address ")
	}
}

func isIPv6(ip string) bool {
	var isIPv6 bool

	switch {
	case strings.Contains(ip, ".") && strings.Contains(ip, ":"):
		isIPv6 = true
	case strings.Contains(ip, "."):
		isIPv6 = false
	case strings.Contains(ip, `:`):
		isIPv6 = true
	}

	return isIPv6
}

func loadIPList(file string) []string {
	ipList, err := os.Open(file)
	// Short way to close the file.
	//defer ipList.Close()

	// Using a closure to actually handle any errors while closing.
	defer func() {
		err = ipList.Close()
		if err != nil {
			log.Fatalf("we're hosed now, jim: %s\n", err)
		}
	}()

	if err != nil {
		log.Fatalf("damnit, jim: %s", err)
	}

	scanner := bufio.NewScanner(ipList)
	scanner.Split(bufio.ScanLines)
	var ipAddresses []string

	for scanner.Scan() {
		ipAddresses = append(ipAddresses, scanner.Text())
	}

	return ipAddresses
}
