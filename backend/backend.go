package backend

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Backend interface {
	SetAlive(bool)
	IsAlive() bool
	GetURL() *url.URL
	GetActiveConnections() int
	GetAvgRespTime() float64
	Serve(http.ResponseWriter, *http.Request)
}

type backend struct {
	url          *url.URL
	alive        bool
	mux          sync.RWMutex
	connections  int
	AvgRespTime  float64
	reverseProxy *httputil.ReverseProxy
}

func (b *backend) GetActiveConnections() int {
	b.mux.RLock()
	connections := b.connections
	b.mux.RUnlock()
	return connections
}

func (b *backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.alive = alive
	b.mux.Unlock()
}

func (b *backend) IsAlive() bool {
	b.mux.RLock()
	alive := b.alive
	defer b.mux.RUnlock()
	return alive
}

func (b *backend) GetURL() *url.URL {
	return b.url
}

func (b *backend) GetAvgRespTime() float64 {
	b.mux.RLock()
	avgRespTime := b.AvgRespTime
	b.mux.RUnlock()
	return avgRespTime
}

func (b *backend) Serve(rw http.ResponseWriter, req *http.Request) {
	defer func() {
		b.mux.Lock()
		b.connections--
		b.mux.Unlock()
	}()

	b.mux.Lock()
	b.connections++
	b.mux.Unlock()
	start := time.Now()
	b.reverseProxy.ServeHTTP(rw, req)
	responseTime := time.Since(start)
	b.mux.Lock()
	alpha := 0.5 // weight for averaging
	newAvg := alpha*float64(responseTime.Milliseconds()) + (1-alpha)*b.AvgRespTime
	b.AvgRespTime = newAvg
	b.mux.Unlock()
}

func NewBackend(u *url.URL, rp *httputil.ReverseProxy) Backend {
	return &backend{
		url:          u,
		alive:        true,
		reverseProxy: rp,
	}
}
