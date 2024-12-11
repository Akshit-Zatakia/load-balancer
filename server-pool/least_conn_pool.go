package serverpool

import (
	"sync"

	"github.com/Akshit-Zatakia/load-balancer/backend"
)

type lcServerPool struct {
	backends []backend.Backend
	mux      sync.RWMutex
}

func (s *lcServerPool) GetNextValidPeer() backend.Backend {
	var leastConnectedPeer backend.Backend
	// Find at least one valid peer
	for _, b := range s.backends {
		if b.IsAlive() {
			leastConnectedPeer = b
			break
		}
	}

	// Check which one has the least number of active connections
	for _, b := range s.backends {
		if !b.IsAlive() {
			continue
		}
		if leastConnectedPeer.GetActiveConnections() > b.GetActiveConnections() {
			leastConnectedPeer = b
		}
	}
	return leastConnectedPeer
}

func (s *lcServerPool) GetBackends() []backend.Backend {
	return s.backends
}

func (s *lcServerPool) AddBackend(b backend.Backend) {
	s.backends = append(s.backends, b)
}

func (s *lcServerPool) GetServerPoolSize() int {
	return len(s.backends)
}
