// Common contains structs, data, and interfaces that
// need to be referenced by other packages
package common

import (
	"time"
)

const (
	// How many bytes of entropy must be produced each entropy cycle
	EntropyVolume int = 32

	// How big a single slice of data is for a host, in bytes
	MinSliceSize int = 512
	MaxSliceSize int = 1048576 // 1 MB

	// How many hosts participate in each quorum
	// This number is chosen to minimize the probability of a single quorum
	// 	becoming more than 80% compromised by an attacker controlling 1/2 of
	// 	the network.
	QuorumSize int = 128

	// How long a single step in the consensus algorithm takes
	StepDuration time.Duration = 15 * time.Second
)

type Entropy [EntropyVolume]byte

// An Address is used to uniquely identify a message recipient.
type Address struct {
	Identifier int
	Host       string
	Port       int
}

// Messages are for sending things over the network.
// Each message has a single destination, and it is
// the job of the network package to interpret the
// destinations.
type Message struct {
	Destination Address
	Payload     []byte
}

type MessageSender interface {
	SendMessage(m Message)
}

type MessageHandler interface {
	HandleMessage(payload []byte)
}

type MessageSender interface {
	Address() Address
	SendMessage(m *Message) error
}
