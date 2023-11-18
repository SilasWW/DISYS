package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	proto "someName/grpc"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	proto.UnimplementedChitChatServer // Necessary
	name                              string
	port                              int
}

var port = flag.Int("port", 0, "server port number")
var lamport int64 = 0
var serverName string = "MAINFRAME"

var clientConns map[int64]proto.ChitChatClient = make(map[int64]proto.ChitChatClient)

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
func (c *Server) Chat(ctx context.Context, in *proto.Publish) (*proto.Broadcast, error) {
	//time handling
	handleLamport(in.ClientLamport)
	
	//send back publish
	lamport++
	return &proto.Broadcast{ServerName: serverName , ServerLamport: lamport, Message: in.Message}, nil
}

//receiving a join message, adding client to list and publish join to all clients
func (c *Server) Join(ctx context.Context, in *proto.Publish) (*proto.Acknowledge, error) {
	//time handling
	handleLamport(in.ClientLamport)

	//dial back to the client and save in map
	lamport++
	dialClient(in.ClientId)

	//broadcast to all clients
	broadMessage(in.ClientId)

	//send back ack
	lamport++
	return &proto.Acknowledge{Name: serverName, ServerLamport: lamport}, nil
}

func dialClient(clientID int64){
	//dial client and connect
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(clientID)), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect to port %d", clientID)
		} else {
			log.Printf("Connected to the Client at port %d\n", clientID)
		}
	
	//add to map
	clientConns[clientID] = proto.NewChitChatClient(conn)
}

func broadMessage(clientID int64){
	var message string = fmt.Sprintf("New person in the chat. Say hi to: %d", clientID)

	for _, value := range clientConns{
		lamport++
		Response, err := value.Trans(context.Background(), &proto.Broadcast{
			ServerName: serverName, Message: message, ServerLamport: lamport,
		})
		if err != nil {
			log.Print(err.Error())
		} else {
			handleLamport(Response.ServerLamport)
			lamport++
		}
	}
	
}
