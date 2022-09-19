package retry

import (
	"time"
)

type Strategy interface {
	Backoff(retries int) time.Duration
}
