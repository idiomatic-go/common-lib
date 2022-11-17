package eventing

type exchange struct {
	msgs map[string]*[]Message
}

func createExchange() *exchange {
	return &exchange{msgs: make(map[string]*[]Message)}
}

func (e *exchange) count(event string) int {
	msgs := e.messages(event)
	if msgs == nil {
		return 0
	}
	return len(msgs)
}

func (e *exchange) add(msg Message) {
	_, ok := e.msgs[msg.Event]
	if !ok {
		var msgs []Message
		msgs = append(msgs, msg)
		e.msgs[msg.Event] = &msgs
		return
	}
	*e.msgs[msg.Event] = append(*e.msgs[msg.Event], msg)
}

func (e *exchange) current(event string) *Message {
	if event == "" {
		return nil
	}
	msgs := e.messages(event)
	if msgs == nil {
		return nil
	}
	len := len(msgs)
	return &msgs[len-1]
}

func (e *exchange) messages(event string) []Message {
	if len(e.msgs) == 0 || event == "" {
		return nil
	}
	if m, ok := e.msgs[event]; ok {
		return *m
	}
	return nil
}
