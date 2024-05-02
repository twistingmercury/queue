package queue_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/twistingmercury/queue"
)

func TestQueue(t *testing.T) {
	q := queue.New[testType]()

	t1 := newTestType()
	t2 := newTestType()
	t3 := newTestType()
	// Test Enqueue
	q.Enqueue(t1)
	q.Enqueue(t2)
	q.Enqueue(t3)

	assert.Equal(t, 3, q.Len(), "Expected queue length to be 3")

	// Test Dequeue
	item, ok := q.Dequeue()
	require.True(t, ok, "Expected Dequeue to return true")
	assert.Equal(t, t1, item, "Expected dequeued item to be 1")
	assert.Equal(t, 2, q.Len(), "Expected queue length to be 2")

	// Test Dequeue on empty queue
	q.Dequeue()
	q.Dequeue()
	item, ok = q.Dequeue()
	assert.False(t, ok, "Expected Dequeue on empty queue to return false")
	assert.Zero(t, item, "Expected Dequeue on empty queue to return zero value")
}

func TestQueueConcurrent(t *testing.T) {
	q := queue.New[testType]()
	var wg sync.WaitGroup

	// Enqueue concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			q.Enqueue(newTestType())
		}(i)
	}
	wg.Wait()

	assert.Equal(t, 1000, q.Len(), "Expected queue length to be 1000")

	// Dequeue concurrently
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			q.Dequeue()
		}()
	}
	wg.Wait()

	assert.Equal(t, 0, q.Len(), "Expected queue to be empty")
}
