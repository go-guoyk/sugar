package sugar_zap

import (
	"go.uber.org/zap"
	"testing"
)

func TestWrap(t *testing.T) {
	zLog, _ := zap.NewDevelopment()
	log := Wrap(zLog)
	log.Info("hello", "name", "test", "age", 17)
}
