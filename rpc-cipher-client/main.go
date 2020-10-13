package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Connecting to cipher.rpc:9000")
	connection, err := grpc.Dial("cipher.rpc:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()
	c := NewCipherServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Encode(ctx, &CipherRequest{Name: "Daher"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response:", r.GetBody())
}
