package pusher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/propsproject/goprops-toolkit/logger"
	"github.com/pusher/pusher-http-go"
	"context"
)

// HandlePrecenseWebHook function to handle member_added & member_removed pusher webhook
func (r *SocketRegistry) HandlePrecenseWebHook(ctx context.Context, req *http.Request) error {
	var payload pusher.Webhook
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ctx.Value("logger").(logger.Wrapper).Error(err)
		return err
	}

	err = json.Unmarshal(b, &payload)
	if err != nil {
		ctx.Value("logger").(logger.Wrapper).Error(err)
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
			registry.RegisterClient(client)
		case "member_removed":
			registry.UnRegisterClient(client)
		}
	}

	return nil
}
