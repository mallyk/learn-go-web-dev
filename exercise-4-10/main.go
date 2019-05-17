package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

		c := make(chan string, 2)

		go serve(conn, c)
		go write(conn, c)
	}
}

func serve(conn net.Conn, c chan<- string) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var i int
		ln := scanner.Text()
		fmt.Println(scanner.Text())
		if i == 0 {
			xs := strings.Fields(ln)
			c <- xs[0]
			c <- xs[1]
		}
		if ln == "" {
			break
		}
		i++
	}
}

func write(conn net.Conn, c <-chan string) {
	defer conn.Close()
	body := "I see you connected with headers"
	body += "\n"
	body += "Method: "
	body += <-c
	body += "\n"
	body += "URI: "
	body += <-c
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
