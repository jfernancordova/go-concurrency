## sync.Mutex

Another way to synchronize access to a shared resource is by using a mutex. A mutex is
named after the concept of mutual exclusion. A mutex is used to create a critical
section around code that ensures only one goroutine at a time can execute that code
section.

## Race Condition

When two or more goroutines have unsynchronized access to a shared resource and
attempt to read and write to that resource at the same time, you have what’s called a
*race condition*. Race conditions are the reason concurrent programming is complicated and has a greater potential for bugs. Read and write operations against a
shared resource must always be atomic, or in other words, done by only one goroutine at a time.

### Bibliography

- [Go in Action](https://www.manning.com/books/go-in-action)
- [A tour of Go](https://go.dev/tour/concurrency/9)

