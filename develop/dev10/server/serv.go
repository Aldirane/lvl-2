package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	stopCh := make(chan interface{})
	listener, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go telnet(conn, stopCh)
		<-stopCh
		close(stopCh)
		break
	}
}

func telnet(conn net.Conn, stopCh chan interface{}) {
	defer conn.Close()
	errCh := make(chan error)

	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				errCh <- err
				return
			}

			log.Print("received: ", data)
			fmt.Fprintf(conn, "response: %s", data)
		}
	}()

	if err := <-errCh; err == io.EOF {
		log.Println("telnet connection dropped")
	} else {
		log.Printf("got error: %v", err)
	}
	stopCh <- 1
}
