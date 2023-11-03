package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "someName/grpc"
	"strconv"
	"time"
)

type Server struct {
	proto.UnimplementedChitChatServer // Necessary
	name                              string
	port                              int
}

var port = flag.Int("port", 0, "server port number")

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

func (c *Server) Chat(ctx context.Context, in *proto.Publish) (*proto.Broadcast, error) {
	return &proto.Broadcast{ServerName: "Only Server", Time: time.Now().String(), Message: in.Message}, nil
}
