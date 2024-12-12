package serverpool

import (
	"sync"

	"github.com/Akshit-Zatakia/load-balancer/backend"
)

type lpServerPool struct {
	backends []backend.Backend
	mux      sync.RWMutex
}

func (s *lpServerPool) GetNextValidPeer() backend.Backend {
	var leastConnectedPeer backend.Backend
	// Find at least one valid peer
	for _, b := range s.backends {
		if b.IsAlive() {
			leastConnectedPeer = b
			break
		}
	}

	// Check which one has the least response time connections
	for _, b := range s.backends {
		if !b.IsAlive() {
			continue
		}
		if b.GetAvgRespTime() < leastConnectedPeer.GetAvgRespTime() {
			leastConnectedPeer = b
		}
	}
	return leastConnectedPeer
}

func (s *lpServerPool) GetBackends() []backend.Backend {
	return s.backends
}

func (s *lpServerPool) AddBackend(b backend.Backend) {
	s.backends = append(s.backends, b)
}

func (s *lpServerPool) GetServerPoolSize() int {
	return len(s.backends)
}
