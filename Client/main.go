package main

import (
	"DISYS_Mock_Exam/gRPC"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var ID string
var serverAddresses = make([]string, 0)
var requestNumber int32 = 0
var messageChannel = make(chan string, 1)
var serverAddressesLock sync.Mutex

func main() {
	ID = uuid.New().String()
	findServerAddresses()
	fmt.Println("Addresses found:")
	for _, address := range serverAddresses {
		fmt.Println(address)
	}
	fmt.Println("Welcome to the system\nEnter an integer and press enter to increment by the value")
	fmt.Println("Notice that each subsequent call must be greater than the previous one")
	fmt.Println("Your username is: ", ID[:3])
	readInputForever()
}

func findServerAddresses() {
	var wg sync.WaitGroup
	wg.Add(90)

	baseString := "localhost:80"
	for i := 10; i < 100; i++ {
		go pingServer(&wg, baseString+strconv.Itoa(i))
	}
	wg.Wait()
}

func pingServer(wg *sync.WaitGroup, address string) {
	log.Println("Pinging server:", address)
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock())
	if err != nil {
		return
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Could not close connection")
		}
	}(conn)

	c := gRPC.NewIncrementSystemClient(conn)

	_, err = c.PingServer(ctx, &gRPC.Ping{ClientID: ID})

	if err == nil {
		addAddressToServerAddresses(address)
	}

	log.Println("Received ping response from:", address)
}

func readInputForever() {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		if number, err := strconv.Atoi(string(line)); err == nil {
			sendIncrementRequestToAll(int32(number))
		} else {
			fmt.Println("Invalid input")
		}
	}

}

func sendIncrementRequest(wg *sync.WaitGroup, address string, value int32) {
	log.Println("Send increment request to:", address)

	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock())
	if err != nil {
		log.Println(address, "did not respond - deleting from server addresses")
		removeAddressFromAddresses(address)
		return
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Could not close connection")
		}
	}(conn)

	c := gRPC.NewIncrementSystemClient(conn)

	response, err := c.Increment(ctx, &gRPC.IncrementRequest{
		Value:     value,
		RequestID: requestNumber,
		ClientID:  ID,
	})

	log.Println("Received response from:", address)

	if err != nil {
		sendToChannelIfNotFull("Your increment is too low\n Previous increment: " + string(response.Value))
	} else {
		sendToChannelIfNotFull("Successful increment\nPrevious value: " + string(response.Value))
	}
}

func sendIncrementRequestToAll(bid int32) {
	requestNumber++
	var wg sync.WaitGroup
	wg.Add(len(serverAddresses))
	for _, serverAddress := range serverAddresses {
		go sendIncrementRequest(&wg, serverAddress, bid)
	}
	wg.Wait()
	fmt.Println(<-messageChannel)
}

func sendToChannelIfNotFull(message string) {
	if len(messageChannel) != 1 {
		messageChannel <- message
	}
}

func removeAddressFromAddresses(address string) {
	serverAddressesLock.Lock()
	defer serverAddressesLock.Unlock()
	index, err := findIndexOfAddress(address)
	if err != nil {
		return
	}
	serverAddresses = append(serverAddresses[:index], serverAddresses[index+1:]...)
}

func findIndexOfAddress(address string) (int, error) {
	for i := 0; i < len(serverAddresses); i++ {
		if serverAddresses[i] == address {
			return i, nil
		}
	}
	return 0, errors.New("could not find address")
}

func addAddressToServerAddresses(address string) {
	serverAddressesLock.Lock()
	defer serverAddressesLock.Unlock()
	serverAddresses = append(serverAddresses, address)
}
