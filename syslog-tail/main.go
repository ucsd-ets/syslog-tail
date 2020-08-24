package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"flag"
	"os"
)

func tailFile(file string) {
	t, err := tail.TailFile(file, tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Printf("%s: %s\n", "SyslogTail", line.Text)
	}
}

var file = flag.String("file", "/var/log/messages", "File to tail, defaults to /var/log/messages")

func main() {
	flag.Parse()

	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Printf("%s does not exist!\n", *file)
		os.Exit(1)
	  }

	fmt.Printf("Tailing file: %s", *file)
	tailFile(*file)
}