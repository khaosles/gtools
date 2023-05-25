package g

import (
	"sync"
)

/*
   @File: number.go
   @Author: khaosles
   @Time: 2023/5/24 17:21
   @Desc: 线程安全的数字
*/

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type SafeNumber[T Numeric] struct {
	value T
	mutex sync.Mutex
}

func NewSafeNumber[T Numeric](i T) SafeNumber[T] {
	return SafeNumber[T]{value: i}
}

func (n *SafeNumber[T]) Inc() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value++
}

func (n *SafeNumber[T]) Dec() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value--
}

func (n *SafeNumber[T]) Add(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value + val
}

func (n *SafeNumber[T]) Sub(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value - val
}

func (n *SafeNumber[T]) Multiply(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value * val
}

func (n *SafeNumber[T]) Divide(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = n.value / val
}

func (n *SafeNumber[T]) Get() T {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	return n.value
}

func (n *SafeNumber[T]) Set(val T) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.value = val
}
