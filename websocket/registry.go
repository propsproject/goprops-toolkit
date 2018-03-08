package websocket

// Registry registry interface to be implemented with whatever tool used for sockets (Pusher, Gorilla, etc)
type Registry interface {
	Run()
	RegisterClient(*Client)
	UnRegisterClient(*Client)
	GetClient(string) (*Client, bool)
}

// Client client interface to be implemented with whatever tool used for sockets (Pusher, Gorilla, etc)
type Client interface{}
