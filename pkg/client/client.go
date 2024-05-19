package client

import (
	"flag"
	"fmt"
	"net"
)

func get_port() int {
    port := flag.Int("port", 8000, "Port to connect to")
    flag.Parse()
    return *port
}

func Client() {
    port := get_port()  
    fmt.Printf("Port: %v\n", port)
    connection, err := net.Dial("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    go handleConnection(connection)
}

func handleConnection(connection net.Conn) {
    fmt.Println("Connected")
    connection.Close()
}
