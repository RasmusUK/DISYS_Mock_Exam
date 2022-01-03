package main

import (
	"DISYS_Mock_Exam/gRPC"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var currentValue int32 = 0
var clientRequestNumber = make(map[string]int32)
var clientPreviousValue = make(map[string]int32)

var clientRequestNumberLock sync.Mutex
var clientPreviousValueLock sync.Mutex
var currentValueLock sync.Mutex

func main() {
	initServer()
	time.Sleep(time.Hour)
}

type server struct {
	gRPC.UnsafeIncrementSystemServer
}

func (s server) Increment(ctx context.Context, request *gRPC.IncrementRequest) (*gRPC.IncrementResponse, error) {
	waitForYourTurn(request.ClientID, request.RequestID)
	setClientRequestNumber(request.ClientID, request.RequestID)
	if request.Value <= clientPreviousValue[request.ClientID] {
		return &gRPC.IncrementResponse{Value: getClientPreviousValue(request.ClientID)}, errors.New("Value is not greater than previous one")
	}
	value := getCurrentValue()
	setClientPreviousValue(request.ClientID, request.Value)
	updateCurrentValue(request.Value)
	return &gRPC.IncrementResponse{Value: value}, nil
}

func (s server) PingServer(ctx context.Context, ping *gRPC.Ping) (*gRPC.Empty, error) {
	log.Println("Received ping from: ", ping.ClientID)
	clientRequestNumber[ping.ClientID] = 0
	clientPreviousValue[ping.ClientID] = 0
	return &gRPC.Empty{}, nil
}

func initServer() {
	var lis net.Listener
	err := errors.New("not initiated yet")
	baseString := "localhost:80"
	counter := 10
	for err != nil && counter < 100 {
		connectionString := baseString + strconv.Itoa(counter)
		log.Println("Trying to start server on port:", connectionString)
		lis, err = net.Listen("tcp", connectionString)
		counter++
	}

	s := grpc.NewServer()
	gRPC.RegisterIncrementSystemServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func waitForYourTurn(clientID string, requestID int32) {
	for {
		if getClientRequestNumber(clientID) == requestID-1 {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func setClientRequestNumber(clientID string, requestID int32) {
	clientRequestNumberLock.Lock()
	defer clientRequestNumberLock.Unlock()
	clientRequestNumber[clientID] = requestID
}

func getClientRequestNumber(clientID string) (requestID int32) {
	clientRequestNumberLock.Lock()
	defer clientRequestNumberLock.Unlock()
	return clientRequestNumber[clientID]
}

func setClientPreviousValue(clientID string, value int32) {
	clientPreviousValueLock.Lock()
	defer clientPreviousValueLock.Unlock()
	clientPreviousValue[clientID] += value
}

func getClientPreviousValue(clientID string) (requestID int32) {
	clientPreviousValueLock.Lock()
	defer clientPreviousValueLock.Unlock()
	return clientPreviousValue[clientID]
}

func updateCurrentValue(value int32) {
	currentValueLock.Lock()
	defer currentValueLock.Unlock()
	currentValue += value
}

func getCurrentValue() (value int32) {
	currentValueLock.Lock()
	defer currentValueLock.Unlock()
	return currentValue
}
