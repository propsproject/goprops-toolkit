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
	Triggers   *sync.Map
}

// Payload ...
type Payload struct {
	ClientID string
	Data     []byte

	//bidirectional channel enables us to "respond" with false to channel sender if processing fails
	Response chan bool
}

// RegisterTriggers add triggers to listen and broadcast for
func (r *SocketRegistry) RegisterTriggers(triggers map[string]Trigger) {
	for key, trigger := range triggers {
		r.Triggers.Store(key, trigger)
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

// NewWorker starts a new worker to listen for each pusher trigger
func (r *SocketRegistry) NewWorker(trigger Trigger) {
	for {
		select {
		case payload := <-trigger.Broadcast:
			if client, ok := r.GetClient(payload.ClientID); ok {
				if ok := client.Send(payload.Data, trigger.Name); !ok {
					payload.Response <- false
				}
			} else {
				payload.Response <- false
			}
		}
	}
}

// DelegatePush send on correct trigger chan
func (r *SocketRegistry) DelegatePush(triggerName string, payload Payload) {
	if t, ok := r.Triggers.Load(triggerName); ok {
		trigger := t.(Trigger)
		trigger.Broadcast <- payload
	} else {
		// trigger not found
		payload.Response <- false
	}
}

// Run registry start a new worker for each trigger and http server for payloads
func (r *SocketRegistry) Run() {
	r.Triggers.Range(func(key, value interface{}) bool {
		go r.NewWorker(value.(Trigger))
		return true
	})
}

// NewPusherRegistry creates a new pusher registry
func NewPusherRegistry(appID, key, secret, cluster string, triggers map[string]Trigger) *SocketRegistry {
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

		registry.RegisterTriggers(triggers)
	})

	return registry
}
