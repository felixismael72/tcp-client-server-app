package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	port := 8000
	server := Server{port}

	server.Start()
}

type Server struct {
	port int
}

func (s Server) Start() {
	dataStream, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatal(err)
	}

	defer dataStream.Close()

	for {
		log.Println("Aguardando...")
		conn, err := dataStream.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.handle(conn)
	}
}

func (s Server) handle(conn net.Conn) {
	defer conn.Close()
	for {
		msg := s.receiveMsg(conn)
		log.Println(msg)
		if msg != "" {
			s.sendBack(conn, []byte(msg))
		} else {
			break
		}
		log.Println(msg)
		break
	}
}

func (s Server) receiveMsg(conn net.Conn) string {
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err == io.EOF {
		return ""
	} else if err != nil {
		log.Fatal(err)
	}
	return data
}

func (s Server) sendBack(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
