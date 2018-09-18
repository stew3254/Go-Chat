package main

import (
        "fmt"
        "net"
        "bufio"
        "os"
)

func receive(conn net.Conn) {
        for {
                msg, err := bufio.NewReader(conn).ReadString('\n')
                if err != nil {
                        fmt.Println(err.Error())
			conn.Close()
			break
                }
                fmt.Print(msg)
        }
}

func main() {
        ipaddr := "0.0.0.0:1234"
        conn, err := net.Dial("tcp", ipaddr)
        if err != nil {
                conn.Close()
        }
        fmt.Println("Server connected")
        for {
                reader := bufio.NewReader(os.Stdin)
                text, _ := reader.ReadString('\n')
                conn.Write([]byte(text))
                go receive(conn)
        }
}
