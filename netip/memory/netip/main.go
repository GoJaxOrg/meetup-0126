// This example demonstrates creating a large amount of IP addresses using the new netip method.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/netip"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run() error {
	ipList := loadIPList("ip-list.txt")
	processIPAddresses(ipList, true)

	return nil
}

func processIPAddresses(ipList []string, print bool) {
	var (
		ipAddresses []netip.Addr
	)
	for _, ip := range ipList {
		tip, _ := netip.ParseAddr(ip)
		ipAddresses = append(ipAddresses, tip)
	}

	if print {
		for i, nip := range ipAddresses {
			fmt.Printf("%d. %s\n", i+1, nip)
		}
	}
}

func loadIPList(file string) []string {
	ipList, err := os.Open(file)
	if err != nil {
		log.Fatalf("damnit, jim: %s", err)
	}
	// Short way to close the file.
	//defer ipList.Close()

	// Using a closure to actually handle any errors while closing.
	defer func() {
		err = ipList.Close()
		if err != nil {
			log.Fatalf("we're hosed now, jim: %s\n", err)
		}
	}()

	scanner := bufio.NewScanner(ipList)
	scanner.Split(bufio.ScanLines)
	var ipAddresses []string

	for scanner.Scan() {
		ipAddresses = append(ipAddresses, scanner.Text())
	}

	return ipAddresses
}
