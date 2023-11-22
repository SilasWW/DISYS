package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	proto "someName/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	id         int
}

var (
	clientPort = flag.Int("cPort", 0, "client port number")
	serverPort = flag.Int("sPort", 0, "server port number")
)
var lamport int64
var serverConnection proto.ChitChatClient

func main() {
	// Parse the flags to get the port for the client
	flag.Parse()

	// Create a client
	client := &Client{
		id: *clientPort,
	}

	//Connect to server
	serverConnection, _ = connectToServer();

	//Send entrance message
	go enterChat(client)

	//Wait for the client (user) to ask for the time
	go waitForMessage(client)

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

func enterChat(client *Client){

	//the following is adapted from grpc.io : https://grpc.io/docs/languages/go/basics/

	//sends initial join to server, receives stream for subscribing to further messages
	lamport++
	stream, err := serverConnection.Join(context.Background(), &proto.Publish{
		ClientId: int64(client.id), Message: "New client joined the server", ClientLamport: lamport,
	})

	// infinite loop receiving new responses from stream
	if err != nil {
			log.Print(err.Error())
		} else {
			log.Print("You joined the chatroom server!")
			for {
				Response,err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil{
					log.Print(err.Error())
				}
				
				handleLamport(Response.ServerLamport)
				lamport++
				
				log.Printf(Response.Message)
			}
		}
			
}

func waitForMessage(client *Client) {

	// Wait for input in the client terminal
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Send Message to the server we get an unimportant response :)
		lamport++
		ReturnMessage, err := serverConnection.Chat(context.Background(), &proto.Publish{
			ClientId: int64(client.id), Message: input, ClientLamport: lamport,
		})

		if err != nil {
			log.Print(err.Error())
		} else {
			handleLamport(ReturnMessage.Lamport)
			lamport++
		}
	}
}