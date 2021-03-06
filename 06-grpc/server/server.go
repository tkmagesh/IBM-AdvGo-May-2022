package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appServiceImpl struct {
	counter int
	proto.UnimplementedAppServiceServer
}

func (asi *appServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	asi.counter++
	time.Sleep(1 * time.Second)
	x := req.GetX()
	y := req.GetY()
	select {
	case <-ctx.Done():
		fmt.Println("Unable to send the response within the given time")
		return nil, errors.New("timeout occured")
	default:
		result := x + y
		fmt.Printf("Received X = %d, Y = %d and returning Result = %d\n", x, y, result)
		res := &proto.AddResponse{
			Result: result,
		}
		return res, nil
	}
}

func (asi *appServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	asi.counter++
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
	asi.counter++
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
	asi.counter++
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
	go func() {
		tick := time.Tick(1 * time.Second)
		for {
			<-tick
			fmt.Printf("Counter : %d\n", asi.counter)
		}
	}()
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
