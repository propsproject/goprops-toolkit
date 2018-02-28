package pusher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// MemberEventPayload either member added or removed
type MemberEventPayload struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
	UserID  string `json:"user_id"`
}

// HandleMemberAdded function to handle member_added pusher webhook
func (r *SocketRegistry) HandleMemberAdded(req *http.Request) error {
	var payload MemberEventPayload
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(b, &payload)
	client := &RegistryClient{
		Registry:    registry,
		ID:          payload.UserID,
		ChannelName: payload.Channel,
	}
	registry.RegisterClient(client)
	return nil
}

// HandleMemberRemoved function to handle member_removed pusher webhook
func (r *SocketRegistry) HandleMemberRemoved(req *http.Request) error {
	var payload MemberEventPayload
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	json.Unmarshal(b, &payload)
	client := &RegistryClient{
		Registry:    registry,
		ID:          payload.UserID,
		ChannelName: payload.Channel,
	}
	registry.UnRegisterClient(client)
	return nil
}
