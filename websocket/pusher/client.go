package pusher

import (
	"encoding/json"

	"github.com/pusher/pusher-http-go"
)

// RegistryClient ...
type RegistryClient struct {
	Registry    *SocketRegistry
	ID          string
	ChannelName string
	PusherConn  *pusher.Client
}

// Send send trigger to pusher
func (c *RegistryClient) Send(data []byte, trigger string) bool {
	var triggerData map[string]interface{}
	json.Unmarshal(data, &triggerData)
	c.PusherConn.Trigger(c.ChannelName, trigger, triggerData)

	return true
}
