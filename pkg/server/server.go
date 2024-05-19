package server

import (
	"flag"
	"fmt"
	"net"
)

func get_port() int {
    port := flag.Int("port", 8000, "port for the server to listen on")
    flag.Parse()
    return *port
}

func Server() {
    // set up port listening
    // upon connection, add client to set of listeners 
    // upon receiving message, publish message to all listeners
    // upon receiving disconnection request, remove client from listeners
    port := get_port()
    incoming := make(chan []byte)
    // outgoing := make(chan []byte)
    connections := make(chan net.Conn)
    fmt.Printf("Port: %v\n", port)
    listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }
    defer listener.Close()

    go getConnection(listener, connections)
    for {
        select {
        case c := <-connections:
            go handleConnection(c)
        case msg := <-incoming:
            go handleIncoming(msg)
        }
    }
}


func getConnection(listener net.Listener, connections chan net.Conn) {
    for {
        connection, err := listener.Accept()
        if err != nil {
                fmt.Printf("%v\n", err)
        }
        connections <- connection
    }
}

func handleConnection(conn net.Conn) {
   fmt.Printf("%v\n", conn.LocalAddr().String()) 
   conn.Close()
}

func handleIncoming(msg []byte) {
    fmt.Printf("%b\n", msg)
}
