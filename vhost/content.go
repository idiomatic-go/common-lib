package vhost

// Credentials function methds
func CreateCredentialsMessage(event, sender string, fn Credentials) Message {
	return CreateMessage(event, sender, fn)
}

func AccessCredentials(msg *Message) Credentials {
	if msg == nil || msg.Content == nil || len(msg.Content) == 0 {
		return nil
	}
	for _, c := range msg.Content {
		fn, ok := c.(Credentials)
		if ok {
			return fn
		}
	}
	return nil
}
