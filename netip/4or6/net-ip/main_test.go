package main

import (
	"testing"
)

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

func TestIsIPv6(t *testing.T) {
	var tests = []struct {
		address string
		desc    string
		want    bool
	}{
		{address: "192.168.0.69", desc: "IPv4", want: false},
		{address: "::ffff:117.188.238.179", desc: "IPv4-mapped IPv6", want: true},
		{address: "::ffff:9c17:d15d", desc: "IPv6", want: true},
	}

	for _, test := range tests {
		t.Log(test.desc)
		got := isIPv6(test.address)
		if got != test.want {
			t.Errorf("want: %t, got: %t\n", test.want, got)
		}
	}
}
