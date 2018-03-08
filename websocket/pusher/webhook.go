package pusher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	lgr "github.com/propsproject/go-utils/logger/v2"
	pusher "github.com/pusher/pusher-http-go"
)

var logger = lgr.Logger

// HandlePrecenseWebHook function to handle member_added & member_removed pusher webhook
func (r *SocketRegistry) HandlePrecenseWebHook(req *http.Request) error {
	var payload pusher.Webhook
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
			ID:          event.UserId,
			ChannelName: event.Channel,
		}
		switch event.Name {
		case "member_added":
			logger.Info("Registering new client to pusher: ", lgr.Field{"ID", client.ID}, lgr.Field{"ChannelName", client.ChannelName})
			registry.RegisterClient(client)
		case "member_removed":
			logger.Info("UnRegistering new client to pusher: ", lgr.Field{"ID", client.ID}, lgr.Field{"ChannelName", client.ChannelName})
			registry.UnRegisterClient(client)
		}
	}

	return nil
}
