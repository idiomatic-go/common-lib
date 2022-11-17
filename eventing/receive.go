package eventing

// Response methods
var resp chan Message

func init() {
	resp = make(chan Message, 100)
	go receive()
}

func receive() {
	for {
		select {
		case msg, open := <-resp:
			// Exit on a closed channel
			if !open {
				return
			}
			Directory.Add(msg.From, msg)
		}
	}
}
