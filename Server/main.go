package main

import (
	"DISYS_Mock_Exam/gRPC"
	"context"
	"errors"
	"fmt"
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

func main() {
	initServer()
	time.Sleep(time.Hour)
}

type server struct {
	gRPC.UnsafeIncrementSystemServer
}

func (s server) Ping(context.Context, *gRPC.Empty) (*gRPC.Empty, error) {
	log.Println("Received ping")
	return &gRPC.Empty{}, nil
}

func (s server) SendBidRequest(_ context.Context, request *gRPC.BidRequest) (*gRPC.BidResponse, error) {
	log.Println("Received bid request from:", request.ClientID[:3])
	if getHighestBid() == 0 {
		auctionIsActive = true
		go auctionTime()
	}
	waitForYourTurn(request.ClientID, request.RequestID)
	setBiddersLogicalTime(request.ClientID, request.RequestID)

	if !auctionIsActive {
		return &gRPC.BidResponse{Success: false}, errors.New("auction is over")
	}

	if setHighestBid(request.Amount) {
		setHighestBidder(request.ClientID)
		return &gRPC.BidResponse{Success: true}, nil
	}

	return &gRPC.BidResponse{Success: false}, nil
}

func (s server) SendResultRequest(_ context.Context, request *gRPC.ResultRequest) (*gRPC.ResultResponse, error) {
	log.Println("Received result request from:", request.ClientID[:3])
	waitForYourTurn(request.ClientID, request.RequestID)
	setBiddersLogicalTime(request.ClientID, request.RequestID)

	if getHighestBid() == 0 {
		return nil, errors.New("no bids has been made")
	}

	name := getHighestBidder()[:3]
	result := "Client " + name + " amount: " + strconv.Itoa(int(getHighestBid()))

	return &gRPC.ResultResponse{
		Result: result,
		Active: auctionIsActive,
	}, nil
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
	gRPC.RegisterBidAuctionClientFEServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func auctionTime() {
	fmt.Println("Auction has started")
	time.Sleep(time.Minute * 1)
	auctionIsActive = false
	fmt.Println("Auction is done")
}

func waitForYourTurn(clientID string, requestID int32) {
	for {
		if getBiddersLogicalTime(clientID) == requestID-1 {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func setBiddersLogicalTime(clientID string, requestID int32) {
	biddersLogicalTimeLock.Lock()
	defer biddersLogicalTimeLock.Unlock()
	clientRequestNumber[clientID] = requestID
}

func getBiddersLogicalTime(clientID string) (requestID int32) {
	biddersLogicalTimeLock.Lock()
	defer biddersLogicalTimeLock.Unlock()
	return clientRequestNumber[clientID]
}

func getHighestBid() (highest int32) {
	highestBidLock.Lock()
	defer highestBidLock.Unlock()
	return highestBid
}

func setHighestBid(input int32) (success bool) {
	highestBidLock.Lock()
	defer highestBidLock.Unlock()
	if input > highestBid {
		highestBid = input
		return true
	}
	return false
}

func getHighestBidder() (name string) {
	highestBidLock.Lock()
	defer highestBidLock.Unlock()
	return highestBidder
}

func setHighestBidder(name string) {
	highestBidderLock.Lock()
	defer highestBidderLock.Unlock()
	highestBidder = name
}
