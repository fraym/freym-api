package client

import (
	"sync"

	"github.com/google/uuid"
)

type requestGate struct {
	mx                     sync.Mutex
	numberOfActiveRequests int
	maxRequests            int
	waitChans              map[string]chan struct{}
}

func newRequestGate(maxRequests int) *requestGate {
	return &requestGate{
		mx:                     sync.Mutex{},
		numberOfActiveRequests: 0,
		maxRequests:            maxRequests,
		waitChans:              map[string]chan struct{}{},
	}
}

func (l *requestGate) Enter() {
	uuid := uuid.New().String()

	for {
		l.mx.Lock()
		if l.numberOfActiveRequests < l.maxRequests {
			if waitChan, ok := l.waitChans[uuid]; ok {
				close(waitChan)
				delete(l.waitChans, uuid)
			}
			l.numberOfActiveRequests++
			l.mx.Unlock()
			return
		}

		var waitChan chan struct{}

		if _, ok := l.waitChans[uuid]; !ok {
			waitChan = make(chan struct{})
			l.waitChans[uuid] = waitChan
		} else {
			waitChan = l.waitChans[uuid]
		}

		l.mx.Unlock()

		<-waitChan
	}
}

func (l *requestGate) Leave() {
	l.mx.Lock()
	defer l.mx.Unlock()

	for _, waitChan := range l.waitChans {
		select {
		case waitChan <- struct{}{}:
		default:
		}
	}

	l.numberOfActiveRequests--
}
