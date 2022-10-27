package startup

import (
	"sync"
	"time"
)

type messageMap map[string]Message

type entry struct {
	uri            string
	c              chan Message
	dependents     []string
	startupStatus  int32
	statusChangeTS time.Time
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

func (s *syncMap) put(e *entry) bool {
	if e == nil || e.uri == "" {
		return false
	}
	s.mu.Lock()
	s.m[e.uri] = e
	s.mu.Unlock()
	return true
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
	e.statusChangeTS = time.Now()
	return true
}

func (s *syncMap) getStatus(uri string) int32 {
	e := s.get(uri)
	if e == nil {
		return StatusEmpty
	}
	return e.startupStatus
}

func (s *syncMap) isStartupSuccessful(uri string) bool {
	status := s.getStatus(uri)
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
	directory = createSyncMap()
}
