package pusher

// Event ...
type Event struct {
	Name      string
	Broadcast chan Payload
}
