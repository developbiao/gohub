package cache

import "time"

type Store interface {
	Set(key string, value string, expireTime time.Duration)
	Get(key string) string
	Has(key string) bool
	Forget(key string)
	Forever(key string, value string)
	Flush()

	IsAlive() error

	// Increment When there is only 1 parameter, its is the key and increases by 1
	// When there are 2 parameters, the fist parameter is the key, and second parameter is the int64 type of
	// the value to be added
	Increment(parameters ...interface{})

	// Decrement When there is only 1 parameter, its is the key and decrement by 1
	// When there are 2 parameters, the fist parameter is the key, and second parameter is the int64 type of
	// the value to be decremented
	Decrement(parameters ...interface{})
}
