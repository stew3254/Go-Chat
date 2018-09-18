package main

import (
        "fmt"
        "net"
        "bufio"
)

func receive(conn net.Conn) {
        for {
                msg, _ := bufio.NewReader(conn).ReadString('\n')
                fmt.Print("[", conn.RemoteAddr(), "] " + msg)
        }
}

func main() {
        service := "0.0.0.0:1234"
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        checkError(err)
        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)
	fmt.Println("Server Listening")
        for {
                conn, err := listener.Accept()
                defer conn.Close()
                fmt.Println("Client [", conn.RemoteAddr(), "] connected")
                if err != nil {
                        continue
                }
                go receive(conn)
        }
}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
        }
}
