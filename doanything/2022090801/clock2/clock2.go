package main

import (
	"fmt"
	"log"
	"net"
	"io"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(fmt.Sprintf("init listen failed: %s", err.Error()))
		return
	}
	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	defer listen.Close()
	for {
		// Accept waits for and returns the next connection to the listener.
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("connet has error: %s", err.Error()))
		}

		go handleConnetion(conn)
	}
}

func handleConnetion(c net.Conn) {
	defer c.Close()
	for {
		// 循环向客户端发送时间信息， 如果出错就结束函数，结束函数前会自动触发c.Close()
		if _, err := io.WriteString(c, time.Now().Format("15:06:07")); err != nil{
			return
		}
	}
}
