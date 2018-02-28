package pusher

import (
	"sync"

	"github.com/matryer/resync"
	"github.com/pusher/pusher-http-go"
)

var registry *SocketRegistry
var onceRegistry resync.Once

// SocketRegistry ...
type SocketRegistry struct {
	PusherConn pusher.Client
	Clients    *sync.Map
	Events     *sync.Map
}

// Payload ...
type Payload struct {
	ClientID string
	Data     []byte

	//bidirectional channel enables us to "respond" with false to channel sender if processing fails
	Response chan bool
}

// RegisterEvents add Events to listen and broadcast for
func (r *SocketRegistry) RegisterEvents(events map[string]Event) {
	for key, event := range events {
		r.Events.Store(key, event)
	}
}

// RegisterClient registers new client with the clients map
func (r *SocketRegistry) RegisterClient(c *RegistryClient) {
	r.Clients.Store(c.ID, c)
}

// UnRegisterClient unregisters a client from the clients map
func (r *SocketRegistry) UnRegisterClient(c *RegistryClient) {
	r.Clients.Delete(c.ID)
}

// GetClient returns the Client pointer by id
func (r *SocketRegistry) GetClient(id string) (*RegistryClient, bool) {
	if c, ok := r.Clients.Load(id); ok {
		return c.(*RegistryClient), true
	}
	return nil, false
}

// NewWorker starts a new worker to listen for each pusher Event
func (r *SocketRegistry) NewWorker(event Event) {
	for {
		select {
		case payload := <-event.Broadcast:
			if client, ok := r.GetClient(payload.ClientID); ok {
				if ok := client.Send(payload.Data, event.Name); !ok {
					payload.Response <- false
				}
			} else {
				payload.Response <- false
			}
		}
	}
}

// DelegateEvent send on correct Event chan
func (r *SocketRegistry) DelegateEvent(eventName string, payload Payload) {
	if t, ok := r.Events.Load(eventName); ok {
		Event := t.(Event)
		Event.Broadcast <- payload
	} else {
		// Event not found
		payload.Response <- false
	}
}

// Run registry start a new worker for each Event and http server for payloads
func (r *SocketRegistry) Run() {
	r.Events.Range(func(key, value interface{}) bool {
		go r.NewWorker(value.(Event))
		return true
	})
}

// NewPusherRegistry creates a new pusher registry
func NewPusherRegistry(appID, key, secret, cluster string, events map[string]Event) *SocketRegistry {
	onceRegistry.Do(func() {
		pusherClient := pusher.Client{
			AppId:   appID,
			Key:     key,
			Secret:  secret,
			Cluster: cluster,
		}

		registry = &SocketRegistry{
			PusherConn: pusherClient,
			Clients:    new(sync.Map),
		}

		registry.RegisterEvents(events)
	})

	return registry
}
