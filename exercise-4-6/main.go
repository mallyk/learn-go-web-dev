package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(scanner.Text())
			if ln == "" {
				break
			}
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	io.WriteString(conn, "I see you connected")
}
