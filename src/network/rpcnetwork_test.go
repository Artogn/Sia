package network

import (
	"common"
	"testing"
)

// a simple message handler
// stores a received message
type TestStoreHandler struct {
	message string
}

func (tsh *TestStoreHandler) StoreMessage(message string, arb *struct{}) error {
	tsh.message = message
	return nil
}

// TestRPCSendMessage tests the NewRPCServer and RegisterHandler functions.
// NewRPCServer must properly initialize a RPC server.
// RegisterHandler must make a RPC available to the client.
// The RPC must sucessfully store a message.
func TestRPCSendMessage(t *testing.T) {
	// create RPCServer and add a message handler
	rpcs, err := NewRPCServer(9988)
	if err != nil {
		t.Fatal("Failed to initialize TCPServer:", err)
	}
	defer rpcs.Close()

	// create message handler and add it to the TCPServer
	tsh := new(TestStoreHandler)
	id := rpcs.RegisterHandler(tsh)

	// send a message
	m := &common.RPCMessage{
		common.Address{id, "localhost", 9988},
		"TestStoreHandler.StoreMessage",
		"hello, world!",
		nil,
	}
	err = SendRPCMessage(m)
	if err != nil {
		t.Fatal("Failed to send message:", err)
	}

	if tsh.message != "hello, world!" {
		t.Fatal("Bad response: expected \"hello, world!\", got \"" + tsh.message + "\"")
	}

	// send a message asynchronously
	tsh.message = ""
	async := SendAsyncRPCMessage(m)
	<-async.Done
	if async.Error != nil {
		t.Fatal("Failed to send message:", async.Error)
	}

	if tsh.message != "hello, world!" {
		t.Fatal("Bad response: expected \"hello, world!\", got \"" + tsh.message + "\"")
	}
}
