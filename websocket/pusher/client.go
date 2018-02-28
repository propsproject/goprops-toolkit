package pusher

import (
	"encoding/json"

	pusher "github.com/pusher/pusher-http-go"
)

// RegistryClient ...
type RegistryClient struct {
	Registry    *SocketRegistry
	ID          string
	ChannelName string
	Pusher      pusher.Client
}

// Send send trigger to pusher
func (c *RegistryClient) Send(data []byte, trigger string) bool {
	var triggerData map[string]interface{}
	json.Unmarshal(data, &triggerData)
	c.Pusher.Trigger(c.ChannelName, trigger, triggerData)

	return true
}
