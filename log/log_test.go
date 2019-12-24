package log

import "testing"

func TestInit(t *testing.T) {
	Init("logs", "kunkka-match", "Kunkka -- ", "info")
	Info("初始化日志: %v\n", "Kunkka日志写入")
}

func BenchmarkInfo(b *testing.B) {
	Init("logs", "kunkka-match", "Kunkka -- ", "info")

	for i := 0; i < b.N; i++ {
		Info("初始化日志: %v\n", "Kunkka日志写入")
	}
}
