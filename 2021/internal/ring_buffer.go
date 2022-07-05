package internal

// RingBuffer so buff. not so thread-safe
type RingBuffer[T any] struct {
	size   int
	head   int
	tail   int
	buffer []T
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		size:   size,
		head:   0,
		tail:   -1,
		buffer: make([]T, 0, size),
	}
}

func (r *RingBuffer[T]) Add(value T) {
	if len(r.buffer) < r.size {
		r.buffer = append(r.buffer, value)
	} else {
		r.buffer[r.head] = value
	}
	r.head = (r.head + 1) % r.size
	r.tail = (r.tail + 1) % r.size
}

func (r *RingBuffer[T]) Values() []T {
	currentSize := len(r.buffer)
	values := make([]T, currentSize, r.size)
	if currentSize == r.size {
		vIndex := 0
		for i := r.head; i != r.tail; i = (i + 1) % r.size {
			values[vIndex] = r.buffer[i]
			vIndex++
		}
		values[vIndex] = r.buffer[r.tail]
	} else {
		copy(values, r.buffer[:r.head])
	}
	return values
}

func (r *RingBuffer[T]) CurrentSize() int {
	return len(r.buffer)
}
