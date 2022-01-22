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
