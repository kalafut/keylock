package main

import (
	"sync"
	"time"
)

type Keylock[T comparable] struct {
	lock       sync.Mutex
	keylocks   map[T]time.Time
	expiration time.Duration
	lastClean  time.Time
}

// NewKeylock returns a new Keylock with the given expiration. If expiration is
// 0, then locks never expire automatically.
func NewKeylock[T comparable](d time.Duration) *Keylock[T] {
	return &Keylock[T]{
		keylocks:   make(map[T]time.Time),
		expiration: d,
	}
}

// Lock attempts to lock the given id. It returns true if the lock was acquired,
// and false otherwise. It does not block.
func (r *Keylock[T]) Lock(id T) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.clean()

	if _, ok := r.keylocks[id]; !ok {
		r.keylocks[id] = time.Now().Add(r.expiration)
		return true
	}
	return false
}

// Unlock unlocks the given id. It is a no-op if the id is not locked.
func (r *Keylock[T]) Unlock(id T) {
	r.lock.Lock()
	defer r.lock.Unlock()

	delete(r.keylocks, id)
}

// clean removes expired locks. It must be called with the lock held.
func (r *Keylock[T]) clean() {
	if len(r.keylocks) == 0 || r.expiration == 0 || time.Since(r.lastClean) < r.expiration {
		return
	}

	now := time.Now()
	for id, t := range r.keylocks {
		if now.After(t) {
			delete(r.keylocks, id)
		}
	}
}
