package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	service := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//doRequestResponse(ctx, service)
	doServerStreaming(ctx, service)
}

func doRequestResponse(ctx context.Context, service proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := service.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Result = ", addResponse.GetResult())
}

func doServerStreaming(ctx context.Context, service proto.AppServiceClient) {
	primeRequest := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := service.GeneratePrimes(ctx, primeRequest)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		primeRes, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime numbers received!")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No %d received\n", primeRes.GetPrimeNo())
	}
}
