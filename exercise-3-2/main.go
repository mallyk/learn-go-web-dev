package main

import (
	"bufio"
	"fmt"
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
			fmt.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	uri := getURI(conn)

	switch uri {
	case "/":
		respond(conn, index())
	case "/test/uri":
		respond(conn, testURI())
	default:
		notFound(conn)
	}
}

func getURI(conn net.Conn) string {
	i := 0
	var uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			uri = strings.Fields(ln)[1]
		}
		if ln == "" {
			break
		}
		i++
	}
	return uri
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			uri := strings.Fields(ln)[1]
			fmt.Println("URI", uri)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func notFound(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>page not found</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 404 Not found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func index() string {
	return `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
}

func testURI() string {
	return `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>you found the test uri</strong></body></html>`
}
