package gorignbuffer

import "fmt"

type RingBuffer[T any] struct {
	buffer   []T
	start    int
	count    int
	capacity int
}

func New[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		panic("ringbuffer: capacity must be > 0")
	}

	return &RingBuffer[T]{
		buffer:   make([]T, capacity),
		capacity: capacity,
	}
}

// Add adds an item to the ring buffer.
// If the buffer is full, it overwrites the oldest item.
func (o *RingBuffer[T]) Add(item T) {
	end := (o.start + o.count) % o.capacity
	o.buffer[end] = item

	if o.count == o.capacity {
		o.start = (o.start + 1) % o.capacity
	} else {
		o.count++
	}
}

// Items returns a slice of items in the ring buffer.
// The items are returned in the order they were added, from oldest to newest.
func (o *RingBuffer[T]) Items() []T {
	result := make([]T, o.count)
	for i := 0; i < o.count; i++ {
		index := (o.start + i) % o.capacity
		result[i] = o.buffer[index]
	}
	return result
}

// Clear resets the ring buffer, removing all items.
// The capacity remains unchanged.
func (o *RingBuffer[T]) Clear() {
	o.start = 0
	o.count = 0
	var zeroValue T
	for i := range o.buffer {
		o.buffer[i] = zeroValue
	}
}

// First returns the first item in the ring buffer.
// It returns an error if the buffer is empty.
func (o *RingBuffer[T]) First() (T, error) {
	if o.count == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("ringbuffer: buffer is empty")
	}

	return o.buffer[o.start], nil
}

// Last returns the last item in the ring buffer.
// It returns an error if the buffer is empty.
func (o *RingBuffer[T]) Last() (T, error) {
	if o.count == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("ringbuffer: buffer is empty")
	}

	end := (o.start + o.count - 1) % o.capacity
	return o.buffer[end], nil
}

// Len returns the number of items in the ring buffer.
// It will never be greater than the capacity.
func (o *RingBuffer[T]) Len() int {
	return o.count
}

// Cap returns the capacity of the ring buffer.
func (o *RingBuffer[T]) Cap() int {
	return o.capacity
}

// Full checks if the ring buffer is full.
func (o *RingBuffer[T]) Full() bool {
	return o.count == o.capacity
}

// Empty checks if the ring buffer is empty.
func (o *RingBuffer[T]) Empty() bool {
	return o.count == 0
}

// Get retrieves an item at the specified index.
// The index is zero-based, where 0 is the oldest item.
func (o *RingBuffer[T]) Get(index int) (T, error) {
	if index < 0 || index >= o.count {
		var zeroValue T
		return zeroValue, fmt.Errorf("ringbuffer: index out of range")
	}
	return o.buffer[(o.start+index)%o.capacity], nil
}

// PopFront removes the first item from the ring buffer.
// It returns the removed item and an error if the buffer is empty.
func (o *RingBuffer[T]) PopFront() (T, error) {
	if o.count == 0 {
		var zeroValue T
		return zeroValue, fmt.Errorf("ringbuffer: buffer is empty")
	}
	item := o.buffer[o.start]
	o.start = (o.start + 1) % o.capacity
	o.count--
	return item, nil
}
