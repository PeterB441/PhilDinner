package main

import (
	"fmt"
)

func main() {
	var channel = make(chan Package)
	go server(channel, 1234)
	go client(channel, 5678)
	for {

	}
}

func server(channel chan Package, seq int) {
	// Try recieve Package to establish connection
	fmt.Println("Server awaiting 1st handshake")
	var message = <-channel
	fmt.Println("Server recieved 1st handshake", "Ack", message.ack, "Seq:", message.seq)
	var seq1st = message.seq + 1 // Includes the plus 1
	message.ack = seq1st
	message.seq = seq
	// Recive Package with seq and ack for the final time
	channel <- message
	fmt.Println("Server sent 3rd handshake")
	message = <-channel
	fmt.Println("Server recieved 3rd handshake", "Ack", message.ack, "Seq:", message.seq)
	if message.ack == seq+1 && message.seq == seq1st {
		fmt.Println("Connection Established")
	} else {
		fmt.Println("Connection Failed", message.ack, message.seq)
	}

}

func client(channel chan Package, seq int) {
	// Try establish connection
	channel <- Package{seq, 0}
	fmt.Println("Client sent 1st handshake")
	// Await acknowledgement package
	var message = <-channel
	fmt.Println("Client recieved 2nd handshake", "Ack", message.ack, "Seq", message.seq)
	if message.ack == seq+1 {
		var ackTemp = message.ack
		message.ack = message.seq + 1
		message.seq = ackTemp
		channel <- message
		fmt.Println("Client sent third handshake")
	}

}

type Package struct {
	seq int
	ack int
}
