package middleware

import (
	"kunkka-match/log"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	log.InitLog()
	Init()
	time.Sleep(3 * time.Second)
}
