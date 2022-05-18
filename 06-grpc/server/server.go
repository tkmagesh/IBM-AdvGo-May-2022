package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()

	result := x + y
	fmt.Printf("Received X = %d, Y = %d and returning Result = %d\n", x, y, result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (asi *appServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			fmt.Printf("Generated Prime # : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			err := serverStream.Send(res)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (asi *appServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	sum := int32(0)
	count := int32(0)
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			avg := sum / count
			res := &proto.AverageResponse{
				Result: avg,
			}
			serverStream.SendAndClose(res)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Printf("Received no for average : %d\n", no)
		sum += no
		count++
	}
	return nil
}

func (asi *appServiceImpl) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		personName := req.GetPerson()
		msg := fmt.Sprintf("Hi %s, %s!", personName.GetFirstName(), personName.GetLastName())
		res := &proto.GreetResponse{
			GreetMessage: msg,
		}
		er := stream.Send(res)
		if er != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func main() {
	asi := &appServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
