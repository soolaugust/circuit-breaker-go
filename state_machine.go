package circuit_breaker_go

import (
	"fmt"
	"log"
	"sync"
)

type Status string
type EventName string
type Handler func() Status

type StateMachine struct {
	mu sync.Mutex
	status Status
	handlers map[Status]map[EventName]Handler
}

// NewStateMachine create state machine with init status
func NewStateMachine(initStatus Status) *StateMachine {
	return &StateMachine{
		status:   initStatus,
		handlers: make(map[Status]map[EventName]Handler),
	}
}

func (sm *StateMachine) getStatus() Status {
	return sm.status
}

func (sm *StateMachine) setStatus(newStatus Status) {
	sm.status = newStatus
}

func (sm *StateMachine) AddHandler(status Status, event EventName, handler Handler) {
	if _, ok := sm.handlers[status]; !ok {
		sm.handlers[status] = make(map[EventName]Handler)
	}
	sm.handlers[status][event] = handler
}

func (sm *StateMachine) Call(event EventName) Status {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	events := sm.handlers[sm.status]
	if events == nil {
		return sm.status
	}
	if fn, ok := events[event]; ok {
		oldStatus := sm.status
		sm.status = fn()
		newStatus := sm.status
		_ = log.Output(1, fmt.Sprintf("status has been changed from %s to %s because of %s\n", oldStatus, newStatus, event))
	}
	return sm.status
}