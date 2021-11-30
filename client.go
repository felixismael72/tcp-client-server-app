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
	host := "localhost"
	port := 8000
	client := Client{host, port}

	for wantToSendMsg() {
		msg := inputMsgToSend()
		response := client.SendMessage(msg)
		fmt.Println("RESPOSTA: ", response)
	}

	fmt.Println("Finalizando...")
}

func inputMsgToSend() string {
	var msg string
	fmt.Print("Digite a mensagem: ")
	if _, err := fmt.Scanln(&msg); err != nil {
		log.Fatal(err)
	}

	return msg
}

func wantToSendMsg() bool {
	var a string

	for {
		fmt.Print("Deseja enviar uma mensagem? (S/N) >> ")
		if _, err := fmt.Scanln(&a); err != nil {
			panic(err)
		}
		a = strings.ToUpper(a)

		if a != "S" && a != "N" {
			fmt.Println("Não entendi!")
		} else {
			break
		}
	}

	return a == "S"
}

type Client struct {
	host string
	port int
}

func (c Client) SendMessage(msg string) string {
	conn := c.getConnection()
	defer conn.Close()

	msg = fmt.Sprintf("%v\n", msg)
	c.sendRequest(conn, []byte(msg))
	return c.getResponse(conn)
}

func (c Client) sendRequest(conn net.Conn, body []byte) {
	if _, err := conn.Write(body); err != nil {
		log.Fatal("Não foi possível enviar o request")
	}
	log.Println("Request enviado")
}

func (c Client) getResponse(conn net.Conn) string {
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err == io.EOF {
		return ""
	} else if err != nil {
		log.Fatal(err)
	}
	return data
}

func (c Client) getConnection() net.Conn {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		log.Fatal("Não foi possível conectar")
	}

	log.Println("Conectando...")
	return conn
}
