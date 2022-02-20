package shift

import (
	"time"
)

type Config struct {
	Network string
	Addr    string
	Timeout *time.Duration
}
