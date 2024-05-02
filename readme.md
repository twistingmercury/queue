# Thread-Safe Queue

This package provides a thread-safe implementation of a queue data structure in Go. The queue allows multiple goroutines to safely enqueue and dequeue elements concurrently without the risk of race conditions or data corruption.

## Why tho?

 had a small REST API that needed some asynchronous processing, but I didn't want to overcomplicate things by introducing a message broker and a separate worker process. The async tasks weren't complex enough to justify the overhead of creating and deploying an entirely new component.

To keep things simple, I came up with a solution where the API would push the tasks onto a queue and then send a signal over a channel of type struct{}. On the other end, a worker goroutine would be listening for these signals. Whenever a signal is received, the worker would dequeue an item from the queue and process it accordingly.

This approach allowed me to implement asynchronous processing within the same codebase, without the need for additional infrastructure or deployments. It struck a good balance between simplicity and effectiveness for my specific use case.

## Features

- Thread-safe concurrent access to the queue
- Generic implementation supporting any element type
- Efficient enqueue and dequeue operations
- Dynamically growing queue size
- Mutex-based synchronization for protecting shared state

## Limitations

If multiple goroutines are concurrently dequeuing items from the queue, the order in which the items are dequeued may not strictly match the order in which they were enqueued. The actual dequeue order will depend on factors such as the timing and scheduling of the goroutines by the Go runtime.

It's important to note that this package does not provide guarantees for maintaining strict ordering of items across multiple goroutines. If your application requires strict ordering of dequeued items, where the order of dequeuing must exactly match the order of enqueuing, even when multiple goroutines are involved, then this package in its current form may not be suitable for your use case.

In such scenarios, you would need to consider alternative synchronization mechanisms or architectures to ensure strict ordering, such as using a single goroutine for dequeuing or implementing additional coordination and communication between goroutines to enforce the desired order.

## Installation

To use the thread-safe queue package in your Go project, you can install it using the following command:

```bash
go get github.com/twistingmercury/queue
```

## Usage

Here's an example of how to use the thread-safe queue:

```go
package main

import (
    "fmt"
    "path/to/queue"
)

func main() {
    q := queue.New[int]()

    // Enqueue elements
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    // Dequeue elements
    item, ok := q.Dequeue()
    if ok {
        fmt.Println("Dequeued item:", item)
    }

    // Get the current length of the queue
    length := q.Len()
    fmt.Println("Queue length:", length)
}
```

In this example, we create a new queue of type `int` using `queue.New[int]()`. We enqueue three elements into the queue using the `Enqueue` method. Then, we dequeue an element using the `Dequeue` method and print its value. Finally, we get the current length of the queue using the `Len` method.

## API

The thread-safe queue provides the following methods:

- `New[T any]() *Queue[T]`: Creates a new empty queue of the specified element type `T`.
- `(q *Queue[T]) Enqueue(item T)`: Enqueues an element `item` into the queue.
- `(q *Queue[T]) Dequeue() (T, bool)`: Dequeues and returns the front element from the queue. Returns the dequeued element and a boolean indicating if the operation was successful.
- `(q *Queue[T]) Len() int`: Returns the current length of the queue.

The queue is implemented using a mutex (`sync.Mutex`) to ensure thread-safe access and prevent race conditions when multiple goroutines access the queue concurrently.

## Performance

The thread-safe queue provides efficient enqueue and dequeue operations with a time complexity of O(1) on average. The mutex-based synchronization introduces some overhead compared to a non-thread-safe implementation, but it ensures the safety and integrity of the queue in concurrent scenarios.

On a MacBook Pro M1Pro the benchmarks produces these results:

```text
go test -bench=. -benchmem -benchtime=10000x -run=^#
goos: darwin
goarch: arm64
pkg: github.com/twistingmercury/queue
BenchmarkQueueEnqueueTestType-10                   10000               151.5 ns/op           426 B/op          0 allocs/op
BenchmarkQueueDequeueTestType-10                   10000                15.14 ns/op            0 B/op          0 allocs/op
BenchmarkQueueConcurrentEnqueueTestType-10         10000               231.5 ns/op           427 B/op          0 allocs/op
BenchmarkQueueConcurrentDequeueTestType-10         10000               123.4 ns/op             0 B/op          0 allocs/op
PASS
```

It's important to consider the trade-offs between thread safety and performance based on the specific requirements of your application. If performance is a critical concern and the queue is accessed frequently, you may need to explore alternative synchronization mechanisms or data structures optimized for concurrent use.

## License

This thread-safe queue package is open-source software licensed under the [MIT License](https://opensource.org/licenses/MIT).