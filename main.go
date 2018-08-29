package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func main() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3015,
	}
	tcp, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer tcp.Close()

	var str string
	for {
		if _, err := fmt.Scanln(&str); err == nil && str != "" && len(str)%2 == 0 {
			src := []byte(str)
			dst := make([]byte, hex.DecodedLen(len(src)))
			if n, err := hex.Decode(dst, src); err != nil {
				fmt.Println(err)
				continue
			} else {
				if _, err := tcp.Write(dst[:n]); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
