package main

import (
	"context"
	"fmt"
	Message "hh/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client() (Message.SendClient, error) {

	conn, err := grpc.Dial("localhost:3333", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := Message.NewSendClient(conn)

	request := &Message.Msg{From: "Thanks for Streaming"}

	for {
		resp, err := client.Sendmsg(context.Background(), request)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive response => %s \n", resp)
		defer conn.Close()
	}
	return client, nil
}

func main() {
	Client()

}
