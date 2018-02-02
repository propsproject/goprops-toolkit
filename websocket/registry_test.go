package websocket

import (
	"log"
	"testing"
)

func TestRegistry(t *testing.T) {
	registry := Registry()
	go registry.Run()

	if registryInit != true {
		t.Error("Expected onceRegistry to be called and set registry init to true")
	}
}

func TestGetClient(t *testing.T) {
	registry := Registry()
	go registry.Run()

	if registryInit != true {
		t.Error("Expected onceRegistry to be called and set registry init to true")
	}

	_, ok := registry.GetClient("id")
	if ok {
		t.Error("Expected registry clients to be empty")
	}
	client := getTestClient()

	registry.RegisterClient(client)

	c, ok := registry.GetClient(client.id)
	if client.id != c.id {
		t.Errorf("Expected client with id: %v got: %v", client.id, c.id)
	}
}

func TestRegisterClient(t *testing.T) {
	registry := Registry()
	go registry.Run()

	client := getTestClient()

	go func(client *Client) {
		registry.RegisterClient(client)
		c, _ := registry.GetClient(client.id)
		if client.id != c.id {
			t.Errorf("Expected client with id: %v got: %v", client.id, c.id)
		}
	}(client)

	registry.UnRegisterClient(client)
	c, ok := registry.GetClient(client.id)
	if ok {
		t.Errorf("Expected false got ok: %v and client id: %v got: %v", ok, client.id, c.id)
	}

}

func TestRegistryBroadcast(t *testing.T) {
	registry := Registry()
	go registry.Run()

	s := getTestWSS()
	defer s.Close()
	id := "PROPS"
	ws, err := getWSConn(id, s.URL)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := []byte("SOME DATA")
	rc := make(chan bool)
	rp := RegistryPayload{expectedData, id, rc}
	registry.Broadcast <- rp

	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Println(err.Error())
	}
	if string(msg) != string(expectedData) {
		t.Fatalf("Expected sent message %v to be %v, message was not echoed back", string(expectedData), string(msg))
	}

	badID := "DOESNT EXIST"
	rp.ClientID = badID
	registry.Broadcast <- rp
	if <-rp.Response != false {
		t.Errorf("Expected false client with id %v does not exist broadcast should send false on response channel", badID)
	}

	defer ws.Close()
}

func getTestClient() *Client {
	var client Client
	client.id = "PROPS"
	client.send = make(chan []byte, 256)
	return &client
}
