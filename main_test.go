package circuit_breaker_go

import (
	"testing"
)

func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker()
	result := cb.Call(ReachFailedThreshold)
	if result != Open {
		t.Fatal("circuit breaker status should be OPEN\n")
	}
	result = cb.Call(Timeout)
	if result != HalfOpen {
		t.Fatal("circuit breaker status should be HALF_OPEN\n")
	}
	result = cb.Call(FailedAgain)
	if result != Open {
		t.Fatal("circuit breaker status should be OPEN\n")
	}
	cb.setStatus(HalfOpen)
	result = cb.Call(Success)
	if result != Close {
		t.Fatal("circuit breaker status should be CLOSE\n")
	}
}
