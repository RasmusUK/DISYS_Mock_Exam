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

func (s server) Increment(_ context.Context, request *gRPC.IncrementRequest) (*gRPC.IncrementResponse, error) {
	log.Println("Received increment from", request.ClientID[:3], "\nClient increment value:", request.Value, "\nClient request number:", request.RequestID)
	waitForYourTurn(request.ClientID, request.RequestID)
	setClientRequestNumber(request.ClientID, request.RequestID)
	if request.Value <= getClientPreviousValue(request.ClientID) {
		return &gRPC.IncrementResponse{Value: 0}, errors.New("value is not greater than previous one")
	}
	value := getCurrentValue()
	setClientPreviousValue(request.ClientID, request.Value)
	updateCurrentValue(request.Value)
	log.Println("Updated value to:", getCurrentValue())
	return &gRPC.IncrementResponse{Value: value}, nil
}

func (s server) PingServer(_ context.Context, ping *gRPC.Ping) (*gRPC.Empty, error) {
	log.Println("Received ping from:", ping.ClientID)
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
		log.Println("Current client request number:", requestID)
		if getClientRequestNumber(clientID) == requestID-1 {
			log.Println("Correct order. Permission granted.")
			break
		}
		log.Println("Wrong request order. Waiting.")
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
	clientPreviousValue[clientID] = value
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
