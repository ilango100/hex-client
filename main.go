package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
)

var p int

func init() {
	flag.IntVar(&p, "p", 3015, "Port number")
}

func main() {
	flag.Parse()

	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: p,
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
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
