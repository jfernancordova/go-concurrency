# Concurrency in Go

![Language](https://img.shields.io/badge/language-Go-orange.svg)&nbsp;

Go has a dedicated concurrency model that makes it particularly useful for applications that require high performance,
scalability, and reliability.

### Rules

- If you don’t need it, don’t use it.
- Don’t communicate by sharing memory, share memory by communicating.
- Don’t over-engineer things by using shared memory and complicated, `error-prone` synchronization primitives,
  instead, use `message-passing` between go routines so variables and data can be used in the appropriate sequence.

### Points

In this project:

- The fundamentals of concurrency in Go - Goroutines
- Techniques that can be used to write concurrent programs
- Concurrency patterns
- Advantages
- Pitfalls

### Sections

#### Goroutines
- [WaitGroups](./waitgroup)

#### Race Conditions
- [Mutexes](./mutex)

#### Channels
- [Producer–consumer problem](./producer_consumer)
- [The Dining Philosophers](./dining_philosophers)
