package internal

import (
	"context"
	"strconv"
	"time"
)

const lastStartTimestampKey = "last_start_timestamp"

// Start records when secure storage started.
func (s *secureStorage) Start(context.Context) error {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	return s.Set(lastStartTimestampKey, timestamp)
}
