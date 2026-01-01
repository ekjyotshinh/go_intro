package examples

import (
	"fmt"
)

// channels are pipes that connect concurrent goroutines
// we can send values into channels from one goroutine and receive those values into another goroutine

func addToChannel(c chan string, val string){
	c <- val
}

func DemoChannels() {

	//channelBasics() // example of basic channel operations

	//channelSynchronization() // example of using channels for synchronization

	//channelDirections() // example of channel directions

	//selectExample() // example of using select with channels

	//nonBlockingSelectExample() // example of non-blocking channel operations using select

	closingChannels() // example of closing channels and receiving from closed channels
}

func channelBasics() {
		// Create a channel
	// channels are unbuffered by default, ie., they only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
	// massive blocking can occur if both the sender and receiver are not ready

	messageChannel := make(chan string)

	// Start a goroutine that sends a message to the channel 
	// if we don't start a separate goroutine here, the program will deadlock, since the send operation will block forever waiting for a receiver
	go addToChannel(messageChannel, "Message from another goroutine!")

	go addToChannel(messageChannel, "Hello, Channels!")

	// Receive the message from the channel 
	msg1 := <-messageChannel
	msg2 := <-messageChannel
	fmt.Println(msg1)
	fmt.Println(msg2)

	// buffered channels
	// we can create buffered channels by specifying a buffer length
	// sends to a buffered channel block only when the buffer is full
	// receives block only when the buffer is empty
	bufferedChannel := make(chan string, 2)

	bufferedChannel <- "Buffered Message 1"
	bufferedChannel <- "Buffered Message 2"

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}
func channelSynchronization() {
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine is doing some work...")
		// simulate work
		done <- true // signal that the work is done
	}()

	<-done // wait for the goroutine to finish
	fmt.Println("Goroutine has completed its work.")
}

func channelDirections() {
	chanExample := make(chan string, 1) // unbuffered channel so it does not block
	sendOnlyChannelExample(chanExample, "valuke sent to channel")
	receivedValue := receiveOnlyChannelExample(chanExample)
	fmt.Println(receivedValue)

}

// Channel directions can be specified in function parameters to restrict the operations that can be performed on the channel
func sendOnlyChannelExample(c chan<- string, val string) {
	fmt.Println("Sending value to channel:", val)
	c <- val // can only send to the channel
}

func receiveOnlyChannelExample(c <-chan string) string {
	fmt.Println("Receiving value from channel")
	return <-c // can only receive from the channel
}

func selectExample() {
	// The select statement lets a goroutine wait on multiple communication operations.
	// A select blocks until one of its cases can run, then it executes that case.
	// It is like a switch statement, but for channels.

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "Message from channel 1"
	}()

	go func() {
		channel2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-channel1:
		fmt.Println("Received:", msg1)
	case msg2 := <-channel2:
		fmt.Println("Received:", msg2)
	}
}

func nonBlockingSelectExample() {
	// basic sends and recieves in channels are blocking operations
	// so we can use select with a default case to implement non-blocking sends, receives, and even non-blocking multi-way selects
	messageChannel := make(chan string)

	// this is a unbuffered channel with no sender yet
	// this receive will not block because of the default case
	// if there were no default this could cause a deadlock
	select {
	case msg := <-messageChannel:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	msgToSend := "Hello, Channel!"

	select {
	case messageChannel <- msgToSend:
		fmt.Println("Sent message:", msgToSend)
	default:
		fmt.Println("No message sent")
	}

	// we can use multiple cases abiuve teh default clause to implement a non-blocking multi-way select
	anotherChannel := make(chan string)

	select {
	case msg := <-messageChannel:
		fmt.Println("Received message from messageChannel:", msg)
	case msg := <-anotherChannel:
		fmt.Println("Received message from anotherChannel:", msg)
	default:
		fmt.Println("No messages received from either channel")
	}

}

func closingChannels() {
	// channels can be closed to indicate that no more values will be sent on them
	// closing a channel is done using the built-in close function
	// only the sender should close a channel, never the receiver
	// closing a channel is useful to communicate completion to the receivers

	dataChannel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			dataChannel <- fmt.Sprintf("Data %d", i)
		}
		close(dataChannel) // close the channel after sending all data
	}()

	// range over the channel to receive values until it's closed
	for data := range dataChannel {
		fmt.Println("Received:", data)
	}

	// use infinite for loop with comma-ok idiom to receive from a closed channel
	anotherDataChannel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			anotherDataChannel <- fmt.Sprintf("Another Data %d", i)
		}
		close(anotherDataChannel)
	}()

	for {
		data, ok := <-anotherDataChannel
		if !ok {
			break // channel is closed, exit the loop
		}
		fmt.Println("Received:", data)
	}

	fmt.Println("Channels closed, no more data.")
}