package main

// AudioHub maintains the set of active clients and broadcasts music packets to the clients
type AudioHub struct {

	// Registered audio clients.
	clients map[*AudioClient]bool

	// Register requests from the audio clients.
	register chan *AudioClient

	// Unregister requests from clients.
	unregister chan *AudioClient
}

func newAudioHub() *AudioHub {
	return &AudioHub{
		register:   make(chan *AudioClient),
		unregister: make(chan *AudioClient),
		clients:    make(map[*AudioClient]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case audioClient := <-h.register:
			h.clients[audioClient] = true
		case audioClient := <-h.unregister:
			if _, ok := h.clients[audioClient]; ok {
				delete(h.clients, audioClient)
				close(audioClient.send)
			}
		}
	}
}
