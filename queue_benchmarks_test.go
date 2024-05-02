package queue_test

import (
	"testing"

	"github.com/twistingmercury/queue"
)

func BenchmarkQueueEnqueueTestType(b *testing.B) {
	q := queue.New[testType]()
	testData := make([]testType, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = newTestType()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(testData[i])
	}
}

func BenchmarkQueueDequeueTestType(b *testing.B) {
	q := queue.New[testType]()
	testData := make([]testType, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = newTestType()
		q.Enqueue(testData[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkQueueConcurrentEnqueueTestType(b *testing.B) {
	q := queue.New[testType]()
	testData := make([]testType, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = newTestType()
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			q.Enqueue(testData[i%b.N])
			i++
		}
	})
}

func BenchmarkQueueConcurrentDequeueTestType(b *testing.B) {
	q := queue.New[testType]()
	testData := make([]testType, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = newTestType()
		q.Enqueue(testData[i])
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Dequeue()
		}
	})
}
