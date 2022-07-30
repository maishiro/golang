package main

import (
	"fmt"
	"net"
)

func main() {
	conn, error := net.Dial("tcp", "localhost:3456")
	if error != nil {
		panic(error)
	}
	defer conn.Close()

	str := "test"
	sendMessage(conn, str)

	str = receiveMessage(conn)
	fmt.Printf("received: [%s]\n", str)
}

func sendMessage(conn net.Conn, str string) {
	_, error := conn.Write([]byte(str))
	if error != nil {
		panic(error)
	}
	fmt.Printf("send: [%s]\n", str)
}

func receiveMessage(conn net.Conn) string {
	var buf = make([]byte, 1024)

	len, error := conn.Read(buf)
	if error != nil {
		panic(error)
	}

	return string(buf[:len])
}
