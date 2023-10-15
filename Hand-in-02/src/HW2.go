package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type packet struct {
	flag string
	seq  int
	ack  int
	id   int
	data string
}

// the data we want to transmit rune for rune
var data = "DISYS"

// 2 channels. c2s for communication from client to server, s2c for communication from server to client
var s2c = make(chan packet)
var c2s = make(chan packet)

var wg sync.WaitGroup

func main() {
	fmt.Println("\n--Starting main--")

	wg.Add(2)

	go client(c2s, s2c, data, len(data))
	go server(c2s, s2c)

	wg.Wait()
}

func client(c2s chan packet, s2c chan packet, datastring string, sendCount int) {
	fmt.Println("--Starting Client--")

	//sends synchronisation to server
	m := packet{flag: "SYN", seq: rand.Intn(1000), ack: 0, data: ""}
	seqCheck := m.seq

	c2s <- m
	fmt.Printf("\nClient sent synchronization: \n Flag: %s \n seq: %d \n ack: %d \n \n", m.flag, m.seq, m.ack)

	//receives synchronisation acknowledment from server, checks acknowledgement value and sets acknowledgement value
	m = <-s2c
	time.Sleep(1)
	if m.flag == "SYN ACK" && m.ack == (seqCheck+1) {
		fmt.Printf("Client received synchronization acknowledgement \n\n")
		m.ack = m.seq + 1
		m.flag = "ACK"
	} else {
		fmt.Println("--ERROR--")
	}

	//sends the data 1 rune at a time with new acknowledgements
	for i := 0; i < sendCount; i++ {
		m.id = i
		m.data = datastring[0 : i+1]

		fmt.Printf("Client sent acknowledgement with data \n Flag: %s \n seq: %d \n ack: %d \n Data: %s \n \n", m.flag, m.seq, m.ack, m.data)

		c2s <- m
		time.Sleep(1)
	}

	//sends FIN message to indicate end of packet transfer, data brings the amount of send packets
	m.flag = "FIN"
	m.data = fmt.Sprint(sendCount)
	c2s <- m

	wg.Done()
}

func server(c2s chan packet, s2c chan packet) {
	fmt.Println("--Starting server--")

	//receives synchronization from client, sets ack value and sends back acknowledgement
	m := <-c2s
	time.Sleep(1)
	seqCheck := 0
	if m.flag == "SYN" {
		fmt.Printf("Server received synchronization \n \n")
		m.ack = m.seq + 1
		m.seq = rand.Intn(1000)
		m.flag = "SYN ACK"
		seqCheck = m.seq
		s2c <- m
		fmt.Printf("Server sent synchronization acknowledgement: \nFlag: %s \n seq: %d \n ack: %d \n \n", m.flag, m.seq, m.ack)
	} else {
		fmt.Println("--ERROR--")
	}

	//creates map for received data and check for accounting flags
	dataPackets := make(map[int]string)
	check := ""

	//receives first data packet header info
	m = <-c2s
	time.Sleep(2)
	if m.flag == "ACK" && m.ack == (seqCheck+1) {
		check = m.flag
	}

	//receives all data packets and adds to map
	for check == "ACK" {
		fmt.Printf("server received acknowledgement \n\n")
		time.Sleep(1)

		datastring := m.data
		id := m.id

		dataPackets[id] = datastring

		m = <-c2s
		check = m.flag
	}

	fmt.Print("all data received \n\n")

	//receives FIN flag and prints out data
	if check == "FIN" {
		countCheck, _ := strconv.Atoi(m.data)

		if countCheck != len(dataPackets) {
			fmt.Printf("--FATAL ERROR PACKAGE LOSS--")
		}

		fmt.Print("--Packets in the order they cam in: \n")
		for id, data := range dataPackets {
			fmt.Println(id, data)
		}

		fmt.Print("\n--Ordered Packets: \n")

		keys := make([]int, 0, len(dataPackets))
		for k := range dataPackets {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			fmt.Println(k, dataPackets[k])
		}
	}

	wg.Done()
}
