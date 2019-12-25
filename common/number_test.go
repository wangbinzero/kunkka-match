package common

import (
	"testing"
	"time"
)

func TestUnwrap(t *testing.T) {
	now := time.Now().UnixNano()
	re := Unwrap(now, 10)
	t.Log(re)
}
