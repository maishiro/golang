package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Socket struct {
	ctx context.Context
}

func NewSocket() *Socket {
	return &Socket{}
}

func (p *Socket) startup(ctx context.Context) {
	p.ctx = ctx
}

func (p *Socket) SetContext(ctx context.Context) {
	p.ctx = ctx

	p.init()
}

func (p *Socket) init() {
	endpoint := ":3456"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", endpoint)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Printf("listening [%s] ...\n", tcpAddr.String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go p.handleClient(conn)
	}
}

func (p *Socket) handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("connected from remote: [%s]\n", conn.RemoteAddr().String())
	rt.EventsEmit(p.ctx, "message", time.Now().Local().Format("2006/01/02 15:04:05.000Z07:00")+fmt.Sprintf("  connected from remote: [%s]\n", conn.RemoteAddr().String()))

	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("received: [%s]\n", string(buf[:len]))

	timestamp := time.Now().Local().Format("2006/01/02 15:04:05.000Z07:00")
	conn.Write([]byte(timestamp))
	fmt.Printf("send: [%s]\n", timestamp)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
