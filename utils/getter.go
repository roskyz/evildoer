package utils

import "sync"

type OnceGetter[T any] struct {
	sync.Once
	fn   func() T
	init T
}

func MakeGetter[T any](fn func() T) OnceGetter[T] {
	return OnceGetter[T]{
		fn: fn,
	}
}

func (og *OnceGetter[T]) Getter() T {
	og.Do(func() {
		og.init = og.fn()
	})
	return og.init
}
