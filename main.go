package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
)

var host string

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "Host string")
}

func main() {
	flag.Parse()

	tcp, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcp.Close()

	var str string
	for {
		if _, err := fmt.Scanln(&str); err == nil && str != "" && len(str)%2 == 0 {
			if dst, err := hex.DecodeString(str); err != nil {
				fmt.Println(err)
				continue
			} else if _, err := tcp.Write(dst); err != nil {
				fmt.Println(err)
			}
		}
	}
}
