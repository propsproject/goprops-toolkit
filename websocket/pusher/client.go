package pusher

import (
	"encoding/json"
)

// RegistryClient ...
type RegistryClient struct {
	Registry    *SocketRegistry
	ID          string
	ChannelName string
}

// Send send trigger to pusher
func (c *RegistryClient) Send(data []byte, trigger string) bool {
	var triggerData map[string]interface{}
	json.Unmarshal(data, &triggerData)
	c.Registry.PusherConn.Trigger(c.ChannelName, trigger, triggerData)

	return true
}
