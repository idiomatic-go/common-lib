package vhost

import "sync"

type list map[string]struct{}

type messageMap map[string]Message

type envelopeMap map[string]Envelope

type entry struct {
	uri           string
	c             chan Message
	dependents    []string
	startupStatus int32
}

type syncMap struct {
	m  map[string]*entry
	mu sync.RWMutex
}

func (s *syncMap) count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func (s *syncMap) data() map[string]*entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m
}

func (s *syncMap) put(uri string, e *entry) {
	s.mu.Lock()
	s.m[uri] = e
	s.mu.Unlock()
}

func (s *syncMap) get(uri string) *entry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[uri]
}

func (s *syncMap) setStatus(uri string, status int32) bool {
	e := s.get(uri)
	if e == nil {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	e.startupStatus = status
	return true
}

func (s *syncMap) getStatus(uri string) (int32, bool) {
	e := s.get(uri)
	if e == nil {
		return StatusEmpty, false
	}
	return e.startupStatus, true
}

func (s *syncMap) isStartupSuccessful(uri string) bool {
	status, _ := s.getStatus(uri)
	return status == StatusSuccessful
}

func (s *syncMap) startupInProgress() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k := range s.m {
		e := s.m[k]
		if e != nil && e.startupStatus == StatusInProgress {
			return k
		}
	}
	return ""
}

func (s *syncMap) startupFailure() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k := range s.m {
		e := s.m[k]
		if e != nil && e.startupStatus == StatusFailure {
			return k
		}
	}
	return ""
}

func empty(m messageMap) {
	if m == nil {
		return
	}
	for k := range m {
		delete(m, k)
	}
}

func createSyncMap() *syncMap {
	return &syncMap{m: make(map[string]*entry)}
}

var directory *syncMap

func init() {
	directory = &syncMap{m: make(map[string]*entry)}
}
