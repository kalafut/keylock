package keylock

import (
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	l := NewLock[string](time.Second)

	if !l.Lock("foo") {
		t.Error("expected lock to succeed")
	}

	if l.Lock("foo") {
		t.Error("expected lock to fail")
	}

	l.Unlock("foo")

	if !l.Lock("foo") {
		t.Error("expected lock to succeed")
	}
}

func TestLockClean(t *testing.T) {
	l := NewLock[string](time.Second)

	if !l.Lock("foo") {
		t.Error("expected lock to succeed")
	}

	time.Sleep(800 * time.Millisecond)

	if l.Lock("foo") {
		t.Error("expected lock to fail")
	}

	if !l.Lock("bar") {
		t.Error("expected lock to succeed")
	}

	time.Sleep(300 * time.Millisecond)

	if !l.Lock("foo") {
		t.Error("expected lock to succeed")
	}

	l2 := NewLock[string](0)
	l2.Lock("foo")
	time.Sleep(1500 * time.Millisecond)
	if l2.Lock("foo") {
		t.Error("expected lock to fail")
	}
}
