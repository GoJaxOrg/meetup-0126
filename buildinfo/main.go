package main

import (
	"debug/buildinfo"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var ver bool
	flag.BoolVar(&ver, "version", false, "Display version/build info.")
	flag.Parse()

	if ver {
		printBuildinfo()
		return
	}

	fmt.Println("woohoo! I didn't manually embed build info!")
}

func printBuildinfo() {
	bin, err := os.Executable()
	if err != nil {
		log.Fatalln("Wasn't quite able to find that binary, bruh.")
	}
	build, err := buildinfo.ReadFile(bin)
	if err != nil {
		log.Fatalln("hit a whoopsie reading your binary")
	}

	fmt.Println("Go Version: ", build.GoVersion)
	fmt.Println("Checksum: ", build.Main.Sum)
	fmt.Println("Dependencies: ", build.Deps)
	for _, val := range build.Settings {
		fmt.Printf("%s: %s\n", val.Key, val.Value)
	}
	fmt.Println("Path: ", build.Path)

}
