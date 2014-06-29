package main

import "crypto/tls"
import "github.com/hydrogen18/test-tls"
import "fmt"
import "net"

func handleClient(conn net.Conn) {
	_, err := fmt.Fprintf(conn, "Hello TLS\n")
	if err != nil {
		fmt.Printf("Error on connection:%v", err)
	}
	conn.Close()

}

func main() {

	config := common.MustGetTlsConfiguration()

	listener, err := tls.Listen("tcp", "localhost:51000", config)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}
		fmt.Printf("Accepted connection\n")
		go handleClient(conn)
	}
}
