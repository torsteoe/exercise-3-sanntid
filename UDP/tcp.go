package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func setUpConnection() {
	addr, err := net.ResolveTCPAddr("tcp", "10.100.23.147:33546")
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		listener, err := net.ListenTCP("tcp", addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	//conn, err := listener.AcceptTCP()
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := make([]byte, 1024)
	_, err = conn.Read(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(msg))
	msgToSend := make([]byte, 1024)
	msgToSend = []byte("Connect to: 10.100.23.187:20008\000")
	_, err = conn.Write(msgToSend)
	if err != nil {
		fmt.Println(err)
		return
	}
	//time.Sleep(2 * time.Second)
}

func server() {
	setUpConnection()

	addr2, err := net.ResolveTCPAddr("tcp", ":20008")
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.ListenTCP("tcp", addr2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dette blir skrevet ut?")
	conn2, err := listener.AcceptTCP()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hei, jeg er forbi accept")
	wg.Add(2)

	go readingTCP(conn2)
	go sendingTCP(conn2)
	wg.Wait()

	/*

		go readingTCP(conn)

		wg.Wait()
	*/
}

func readingTCP(conn *net.TCPConn) {
	defer wg.Done()
	for {
		msg := make([]byte, 1024)
		fmt.Println("lese melding")
		_, err := conn.Read(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
	}
}

func sendingTCP(conn *net.TCPConn) {
	defer wg.Done()
	for {
		msgToSend := make([]byte, 1024)
		fmt.Println("Skrive melding")
		msgToSend = []byte("Gratulerer\x00")
		_, err := conn.Write(msgToSend)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	time.Sleep(2)
	fmt.Println("tull")

	server()

}
