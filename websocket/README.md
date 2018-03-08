# websocket

## Overview

The websocket package contains a set of interfaces and packages that implement those interfaces using whatever websocket/pub-sub/queuing technology
needed behind the scenes. For example look at the pusher package. This package satifies the websocket packages but uses the Pusher API for messaging under the hood. Every package implemented will have the same exact architecture. At a high level to manage sockets, connections and updates we make use of a Registry and a Client struct. The Registry is used to keep track of multiple clients connecting to the backend, properly identify the correct client connection, and broad casting messages to proper channels, similar to a traffic cop. Client structs hold the active bi-directional connection to the "other-side", unique identifier, and a send channel that knows how to properly write to that connection. The client is also responsible for keepalive strategies.

### Click Testing

Using the pusher package as an example, it is pretty easy to simulate a near real environment without having to have your client application running. We use ngrok to tunnel http request to our localhost machines, this allows us to handle the pusher webhooks and see whats happening locally in real time adding faster development time, ngrok also provides the ability to replay request. Next all we needed was simple HTML and JS, using JSBin we were able to make pusher do "stuff" and make sure our backend code is functioning as intended.

- [https://ngrok.com/](ngrok)
- [https://gist.github.com/kc1116/6e650f0ac5b7e4ce5fbd29e3bdeb22a1](gist) use for jsbin
- [https://pusher.com/docs](pusher)