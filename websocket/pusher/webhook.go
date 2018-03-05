package pusher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// WebhookPayload ...
type WebhookPayload struct {
	TimeMS time.Time      `json:"time_ms"`
	Events []EventPayload `json:"events"`
}

// EventPayload either member added or removed
type EventPayload struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
	UserID  string `json:"user_id"`
}

// HandlePrecenseWebHook function to handle member_added & member_removed pusher webhook
func (r *SocketRegistry) HandlePrecenseWebHook(req *http.Request) error {
	var payload WebhookPayload
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(b, &payload)

	for _, event := range payload.Events {
		client := &RegistryClient{
			Registry:    registry,
			ID:          event.UserID,
			ChannelName: event.Channel,
		}
		switch event.Name {
		case "member_added":
			registry.RegisterClient(client)
		default:
			registry.UnRegisterClient(client)
		}
	}

	return nil
}
