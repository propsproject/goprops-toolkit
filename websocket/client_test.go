package websocket

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestRegisterNewClient(t *testing.T) {
	registry := Registry()
	go registry.Run()

	s := getTestWSS()
	defer s.Close()
	id := "PROPS"
	ws, err := getWSConn(id, s.URL)
	if err != nil {
		t.Fatal(err)
	}

	c, ok := registry.GetClient(id)
	if id != c.id {
		t.Errorf("Expected client with id: %v got: %v ok: %v", id, c.id, ok)
	}

	defer ws.Close()
}

func TestClientPing(t *testing.T) {
	registry := Registry()
	go registry.Run()

	s := getTestWSS()
	defer s.Close()
	id := "PROPS"
	ws, err := getWSConn(id, s.URL)
	if err != nil {
		t.Fatal(err)
	}

	ws.Close()

	//give client sometime to discover ws conn is dead and clean its self up
	time.Sleep(time.Second * 40)

	c, ok := registry.GetClient(id)
	if ok {
		t.Errorf("Expected client to clean it's self up and unregister instead got ok: %v and client with id: %v", ok, c.id)
	}
}

func getTestWSS() *httptest.Server {
	// simulates a websocket server request coming from client side app
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RegisterNewClient(registry, r.Header.Get("id"), w, r)
	}))

	return s
}

func getWSConn(id, sURL string) (*websocket.Conn, error) {
	url := "ws" + strings.TrimPrefix(sURL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(url, http.Header{"id": []string{id}})
	if err != nil {
		return nil, err
	}

	return ws, nil
}
