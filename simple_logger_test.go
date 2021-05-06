package sugar

import (
	"errors"
	"testing"
)

func TestSimpleLogger(t *testing.T) {
	s := NewSimpleLogger(true)
	s.Error("hello", "world", 1, true, errors.New("baka"))
}
