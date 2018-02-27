package websocket

import (
	"net/http"
	"sync"

	"github.com/matryer/resync"
	"github.com/propsproject/go-utils/logger"
)

var registry *SocketRegistry
var onceRegistry resync.Once

var registryInit bool

//SocketRegistry ...
type SocketRegistry struct {
	// Registered clients.
	clients *sync.Map

	// Inbound messages from the clients.
	Broadcast chan RegistryPayload
}

//RegistryPayload data is the raw data to be sent on the socket, client id is the id of the client and response is a channel to communicate client not found
type RegistryPayload struct {
	Data     []byte    `json:"data"`
	ClientID string    `json:"clientid"`
	Response chan bool `json:"response"`
}

type getRequest struct {
	ch chan *Client
	id string
}

//Registry creates a new websocket registry
func Registry() *SocketRegistry {
	onceRegistry.Do(func() {
		registry = &SocketRegistry{
			Broadcast: make(chan RegistryPayload),
			clients:   new(sync.Map),
		}
	})

	registryInit = true
	return registry
}

//GetClient returns the Client pointer by id
func (r *SocketRegistry) GetClient(id string) (*Client, bool) {
	if c, ok := r.clients.Load(id); ok {
		return c.(*Client), true
	}
	return &Client{}, false
}

//RegisterClient registers new client with the clients map
func (r *SocketRegistry) RegisterClient(c *Client) {
	r.clients.Store(c.id, c)
}

//UnRegisterClient unregisters a client from the clients map
func (r *SocketRegistry) UnRegisterClient(c *Client) {
	r.clients.Delete(c.id)
	if c.send != nil {
		close(c.send)
	}
}

//Run start listening for messages to broadcast to clients
func (r *SocketRegistry) Run() {
	for {
		select {
		case payload := <-r.Broadcast:
			if client, ok := r.GetClient(payload.ClientID); ok {
				select {
				case client.send <- payload.Data:
				default:
					r.UnRegisterClient(client)
				}
			} else {
				payload.Response <- false
			}
		}
	}
}

//RegisterNewClient handles websocket requests from the peer.
func (r *SocketRegistry) RegisterNewClient(registry *SocketRegistry, id string, w http.ResponseWriter, req *http.Request) error {

	if _, ok := r.GetClient(id); !ok {
		conn, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			return err
		}
		client := &Client{registry: registry, conn: conn, send: make(chan []byte, 256), id: id}
		client.registry.RegisterClient(client)

		go client.writePump()
		go client.handleSigTerm()
		return nil
	}

	logger.Info("Client already registered")
	return nil
}
