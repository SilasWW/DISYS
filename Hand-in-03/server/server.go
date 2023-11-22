package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	proto "someName/grpc"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChitChatServer // Necessary
	name                              string
	port                              int
}

var port = flag.Int("port", 0, "server port number")
var lamport int64 = 0
var serverName string = "MAINFRAME"
var mList []string

func main() {
	// Get the port from the command line when the server is run
	flag.Parse()

	// Create a server struct
	server := &Server{
		name: "serverName",
		port: *port,
	}

	// Start the server
	go startServer(server)

	// Keep the server running until it is manually quit
	for {

	}

}

func handleLamport(clientLamport int64){
	if(clientLamport > lamport){
		lamport = clientLamport
	}
}

func startServer(server *Server) {

	// Create a new grpc server
	grpcServer := grpc.NewServer()

	// Make the server listen at the given port (convert int port to string)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d\n", server.port)

	// Register the grpc server and serve its listener
	proto.RegisterChitChatServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}
}

//receiving a chat message and then broadcasting it to all clients
func (c *Server) Chat(ctx context.Context, in *proto.Publish) (*proto.Acknowledge, error) {
	//time handling
	handleLamport(in.ClientLamport)

	//add message to list to be broadcasted
	var message string = fmt.Sprintf("%d: %s",in.ClientId, in.Message)
	mList = append(mList, message)

	//send back ack
	lamport++
	return &proto.Acknowledge{Name: serverName , Lamport: lamport}, nil
}

//receiving a join message, starting async function that streams messages to client (subscription)
func (c *Server) Join(in *proto.Publish, stream proto.ChitChat_JoinServer) error {
	//time handling
	handleLamport(in.ClientLamport)

	//add new welcome message to list 
	var message string = fmt.Sprintf("%d : %s", in.ClientId, in.Message)
	mList = append(mList,message)

	//subscribe client to messages
	//unending loop checking lenght of slice 
	var messageKnown int = len(mList);
	for {
		if messageKnown < len(mList) {
			for _, v := range mList{
				if err := stream.Send(&proto.Broadcast{
					ServerName: serverName, Message: v, ServerLamport: lamport,
				}); err != nil{
					return err
				}
			}
			messageKnown = len(mList)
		}
		time.Sleep(1 * time.Second)
	}
}

