package main

import (
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func TestProcessIPAddresses(t *testing.T) {
	var ipList = []string{
		"255.245.217.174",
		"156.23.209.93",
		"144.33.87.231",
		"82.209.54.189",
		"16.205.51.164",
	}
	processIPAddresses(ipList, true)
}

func TestProcess1KIPAddresses(t *testing.T) {
	ipList := loadIPList("ip-list.txt")
	processIPAddresses(ipList, false)
}

func TestProcess10KIPAddresses(t *testing.T) {
	ipList := loadIPList("10k-list.txt")
	processIPAddresses(ipList, false)
}

func TestLoadFile(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{file: "ip-list.txt", want: 1000},
		{file: "10k-list.txt", want: 10000},
	}

	for _, test := range tests {
		list := loadIPList(test.file)
		got := len(list)
		if got != test.want {
			t.Errorf("want: %d, got: %d\n", test.want, got)
		}
	}
}

func BenchmarkProcessAllIPAddresses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ipList := loadIPList("10k-list.txt")
		processIPAddresses(ipList, false)
	}
	writeHeapProfile()
}

func TestProfileProcess10KAddresses(t *testing.T) {
	t.Log("TestProfileProcess10KAddresses")

	ipList := loadIPList("10k-list.txt")
	processIPAddresses(ipList, true)
	writeHeapProfile()
}

func writeHeapProfile() {
	outFile, err := os.Create("profile.mem")
	if err != nil {
		log.Fatalf("error creating profile destination: %s\n", err)
	}

	err = pprof.WriteHeapProfile(outFile)
	if err != nil {
		log.Fatalf("error writing heap output: %s\n", err)
	}

	err = outFile.Close()
}
