package examples

import "fmt"

// enum are special case of iota where we use it to create a set of related constants.

// this enum represents the states of a server and has an underlying int type.
type ServerState int

const (
	StateIdle ServerState = iota	//0
	StateStarting					//1
	StateRunning					//2
	StateStopping					//3
	StateStopped					//4
)

var stateName = map[ServerState]string{
	StateIdle:     "Idle",
	StateStarting: "Starting",
	StateRunning:  "Running",
	StateStopping: "Stopping",
	StateStopped:  "Stopped",
}

// String method to get the string representation of the ServerState enum using the map created, otherwise they would shouw as integers.
func (s ServerState) String() string {
	if name, ok := stateName[s]; ok {
		return name
	}
	return "Unknown"
}
func DemoEnum() {
	var state ServerState = StateStarting
	fmt.Println("Server state:", state)

	state = StateRunning
	fmt.Println("Server state:", state)

	state = StateStopping
	fmt.Println("Server state:", state)
}