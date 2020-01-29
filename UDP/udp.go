package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func server() {
	udpResolve, err := net.ResolveUDPAddr("udp", ":20008")

	if err != nil {
		fmt.Println(err)
		return
	}
	connection, err := net.ListenUDP("udp", udpResolve)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	udpResolve2, err := net.ResolveUDPAddr("udp", "10.100.23.147:20008")

	if err != nil {
		fmt.Println(err)
		return
	}
	//connection2, err := net.ListenUDP("udp", udpResolve2)

	if err != nil {
		fmt.Println(err)
		return
	}
	//defer connection2.Close()
	wg.Add(2)
	go reading(connection)
	go sending(connection, udpResolve2)
	wg.Wait()
}

func reading(conn *net.UDPConn) {
	defer wg.Done()
	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client ", addr)
		fmt.Println("Bytes = ", n)
		fmt.Println(string(buf))
	}

}

func sending(conn *net.UDPConn, udpAddr *net.UDPAddr) {
	defer wg.Done()
	for {
		msgToClient := []byte("Hello server ")
		time.Sleep(2 * time.Second)
		//	msgToClient := []byte("Client said: " + string(buf))

		_, err := conn.WriteToUDP(msgToClient, udpAddr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
func main() {
	server()
	fmt.Println("hei på deg")
}
