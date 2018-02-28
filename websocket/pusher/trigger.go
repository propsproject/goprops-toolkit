package pusher

// Trigger create trigger to listen for Payload and send correct pusher trigger
type Trigger struct {
	Name      string
	Broadcast chan Payload
}
