# keylock

Keylock lets your application control exclusive access to many keys. Unlike a `sync.Lock`, it does not block, but instead returns `true` or `false` depending on whether the lock was acquired. Locks can be cheaply acquired for many keys without create per-key `sync.Mutex`es. An additional difference it that a lock will be released automatically after a user-defined timeout, if they've not already been released manually.

At < 60 lines, you may find it easier to just copy the code into your project than to take the dependency. I've found it useful in a few projects, so I'm publishing it here.

## Usage

```go

// Locks can be keyed with any comparable type. The timeout can be set to 0 to
// disable automatic unlocking.
l := NewKeylock[string](time.Second)

l.Lock("foo")      // true
l.Lock("foo")      // false
l.Lock("bar")      // true
l.Unlock("foo")    // Note: this is safe to call even if the lock is not held.

time.Sleep(1100*time.Millisecond)
l.Lock("bar")      // true (the lock was released after 1 second)
```
