package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	dataStream, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer dataStream.Close()

	for {
		conn, err := dataStream.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
	}

	conn.Close()
}
