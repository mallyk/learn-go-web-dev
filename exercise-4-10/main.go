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

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var xs []string
	var i int
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(scanner.Text())
		if i == 0 {
			xs = strings.Fields(ln)
		}
		if ln == "" {
			break
		}
		i++
	}

	write(conn, xs[0], xs[1])
}

func write(conn net.Conn, method string, uri string) {
	defer conn.Close()
	body := "I see you connected with headers"
	body += "\n"
	body += "Method: "
	body += method
	body += "\n"
	body += "URI: "
	body += uri
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
