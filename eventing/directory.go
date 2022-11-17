package eventing

import (
	"fmt"
	"sync"
)

type entry struct {
	uri  string
	c    chan Message
	msgs *exchange
}

type syncMap struct {
	m  map[string]*entry
	mu sync.RWMutex
}

var Directory *syncMap

func createSyncMap() *syncMap {
	return &syncMap{m: make(map[string]*entry)}
}

func createEntry(uri string, c chan Message) *entry {
	return &entry{uri: uri, c: c, msgs: createExchange()}
}

func init() {
	Directory = createSyncMap()
}

func (s *syncMap) Exists(uri string) bool {
	if uri == "" {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.m[uri]; ok {
		return true
	}
	return false
}

func (s *syncMap) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func (s *syncMap) CurrentMessage(uri, event string) *Message {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.m[uri]
	if e == nil || e.msgs == nil {
		return nil
	}
	return e.msgs.current(event)
}

func (s *syncMap) Broadcast(event string, status int32) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	msg := CreateMessage("", VirtualHost, event, status, nil)
	for key, entry := range s.m {
		msg.To = key
		entry.c <- msg
	}
}

func (s *syncMap) FindStatus(event string, status int32) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k := range s.m {
		e := s.m[k]
		if e == nil {
			continue
		}
		msg := e.msgs.current(event)
		if msg != nil && msg.Status == status {
			return k
		}
	}
	return ""
}

func (s *syncMap) SendMessage(msg Message) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.m[msg.To]
	if e == nil {
		return fmt.Errorf("invalid argument : to uri invalid %v", msg.To)
	}
	if e.c == nil {
		return fmt.Errorf("invalid initialization : channel is nil %v", msg.To)
	}
	e.c <- msg
	return nil
}

func (s *syncMap) Uri() []string {
	var uri []string
	s.mu.RLock()
	defer s.mu.RUnlock()
	for key, _ := range s.m {
		uri = append(uri, key)
	}
	return uri
}

func (s *syncMap) get(uri string) *entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[uri]
}

func (s *syncMap) put(e *entry) bool {
	if e == nil || e.uri == "" {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[e.uri] = e
	return true
}

func (s *syncMap) Put(uri string, c chan Message) bool {
	if uri == "" {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	e := createEntry(uri, c)
	s.m[e.uri] = e
	return true
}

func (s *syncMap) GetStatus(uri, event string) int32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.m[uri]
	if e == nil {
		return 0
	}
	msg := e.msgs.current(event)
	if msg == nil {
		return 0
	}
	return msg.Status
}

func (s *syncMap) Add(uri string, msg Message) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	e := s.m[uri]
	if e == nil {
		return false
	}
	e.msgs.add(msg)
	return true
}

func (s *syncMap) AddStatus(uri, event string, status int32) bool {
	return s.Add(uri, CreateMessage(VirtualHost, VirtualHost, event, status, nil))
}

func (s *syncMap) Shutdown() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, entry := range s.m {
		if entry.c != nil {
			close(entry.c)
		}
		delete(s.m, key)
	}
}
