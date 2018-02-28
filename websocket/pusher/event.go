package pusher

// Event create Event to listen for Payload and send correct pusher Event
type Event struct {
	Name      string
	Broadcast chan Payload
}
