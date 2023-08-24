package main

import (
	"context"
	"fmt"
	Message "hh/proto"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	defaulHost  = "localhost"
	defaultPort = "3333"
)

func Client() (Message.SendClient, error) {
	if os.Getenv("SERVICE_HOST") != "" {
		defaulHost = os.Getenv("SERVICE_HOST")
	}
	if string(os.Getenv("SERVICE_PORT")) != "" {
		defaultPort = string(os.Getenv("SERVICE_PORT"))
	}
	log.Println(defaulHost)
	addr := fmt.Sprintf("%s:%s", defaulHost, defaultPort)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := Message.NewSendClient(conn)

	request := &Message.Msg{From: "Thanks for Streaming"}

	stream, err := client.Sendmsg(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(feature)
	}

	return client, nil
}

func main() {
	Client()
}
