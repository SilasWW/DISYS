package main

import (
	"bufio"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	proto "someName/grpc"
	"strconv"
)

type Client struct {
	id         int
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number")
)
var lamport int64

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id: *clientPort,
	}

	// Wait for the client (user) to ask for the time
	go waitForTimeRequest(client)

	for {

	}
}

func handleLamport(serverLamport int64){
	if(serverLamport > lamport){
		lamport = serverLamport
	}
}

func connectToServer() (proto.ChitChatClient, error) {
	// Dial the server at the specified port.
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(*serverPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *serverPort)
	} else {
		log.Printf("Connected to the server at port %d\n", *serverPort)
	}
	return proto.NewChitChatClient(conn), nil
}

func waitForTimeRequest(client *Client) {
	// Connect to the server
	serverConnection, _ := connectToServer()

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		log.Printf("You send the message: %s to the server \n", input)

		// Ask the server for the time
		lamport++
		ReturnMessage, err := serverConnection.Chat(context.Background(), &proto.Publish{
			ClientId: int64(client.id), Message: input, ClientLamport: lamport,
		})

		if err != nil {
			log.Print(err.Error())
		} else {
			handleLamport(ReturnMessage.ServerLamport)
			log.Printf("MAINFRAME: New message: ' %s ' at lamport time: %d \n", ReturnMessage.Message, lamport)
		}
	}
}
