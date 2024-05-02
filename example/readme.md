# Thread-Safe Queue Example

This example demonstrates how to use the thread-safe queue package to enqueue and dequeue items concurrently using multiple goroutines.

## Prerequisites

- Go programming language installed (version 1.21 or later)
- Thread-safe queue package installed (`github.com/twistingmercury/queue`)

## Running the Example

1. Make sure you have the thread-safe queue package installed. You can install it using the following command:

   ```bash
   go get -u github.com/twistingmercury/queue
   ```

3. Run the example using the following command:

   ```bash
    ➜  queue cd example 
    ➜  example go run example.go 
   ```

## Example Output

When you run the example, you will see output similar to the following:

```
Dequeued item: 1
Dequeued item: 2
Dequeued item: 3
Dequeued item: 4
Dequeued item: 5
Dequeued item: 6
Dequeued item: 7
Dequeued item: 8
Dequeued item: 9
Dequeued item: 10
All items dequeued
```

The output shows the dequeued items from the queue. However, the order of dequeuing may not necessarily match the order of enqueuing due to the concurrent nature of the goroutines.
