package pusher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lgr "github.com/propsproject/go-utils/logger/v2"
)

var logger = lgr.Logger

// WebhookPayload ...
type WebhookPayload struct {
	TimeMS int64          `json:"time_ms"`
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
		logger.Error(err)
		return err
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, event := range payload.Events {
		client := &RegistryClient{
			Registry:    registry,
			ID:          event.UserID,
			ChannelName: event.Channel,
		}
		switch event.Name {
		case "member_added":
			logger.Info("Registering new client to pusher: ", lgr.Field{"ID", client.ID}, lgr.Field{"ChannelName", client.ChannelName})
			registry.RegisterClient(client)
		default:
			logger.Info("UnRegistering new client to pusher: ", lgr.Field{"ID", client.ID}, lgr.Field{"ChannelName", client.ChannelName})
			registry.UnRegisterClient(client)
		}
	}

	return nil
}
