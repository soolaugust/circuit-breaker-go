package circuit_breaker_go

var (
	// Open circuit breaker is open
	// request will get fatal response immediately
	Open = Status("OPEN")
	// HalfOpen circuit breaker is waiting
	// if next requests success, will change to CLOSE
	// otherwise, will change to OPEN
	HalfOpen = Status("HALF_OPEN")
	// Close circuit breaker is closed
	// request will get response from remote service
	Close = Status("CLOSE")
)

var (
	// ReachFailedThreshold will change circuit breaker to OPEN
	ReachFailedThreshold = EventName("FAILED_THRESHOLD_REACHED")
	// Timeout will change circuit breaker to HALF_OPEN
	Timeout = EventName("TIMEOUT")
	// FailedAgain will change circuit breaker to OPEN
	FailedAgain = EventName("FAILED_AGAIN")
	// Success will change circuit breaker to CLOSE
	Success = EventName("SUCCESS")
)

type CircuitBreaker struct {
	*StateMachine
}

func NewCircuitBreaker() *CircuitBreaker {
	// define a state machine
	cb := &CircuitBreaker{
		StateMachine: NewStateMachine(Close),
	}
	// CLOSE -> OPEN : failed count reached threshold
	cb.AddHandler(Close, ReachFailedThreshold, cb.ReachFailedThreshold)
	// OPEN -> HALF_CLOSE : timeout
	cb.AddHandler(Open, Timeout, cb.Timeout)
	// HALF_OPEN -> OPEN : failed again
	cb.AddHandler(HalfOpen, FailedAgain, cb.FailedAgain)
	// HALF_OPEN -> CLOSE : success
	cb.AddHandler(HalfOpen, Success, cb.Success)
	return cb
}

func (cb *CircuitBreaker) ReachFailedThreshold() Status {
	return Open
}

func (cb *CircuitBreaker) Success() Status {
	return Close
}

func (cb *CircuitBreaker) Timeout() Status {
	return HalfOpen
}

func (cb *CircuitBreaker) FailedAgain() Status {
	return Open
}