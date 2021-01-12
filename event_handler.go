package circuit_breaker_go

type EventHandler interface {
	ReachFailedThreshold() Status
	Success() Status
	Timeout() Status
	FailedAgain() Status
}
