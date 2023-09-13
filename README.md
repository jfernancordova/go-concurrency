## Concurrency in Go

![Language](https://img.shields.io/badge/language-Go-orange.svg)&nbsp;

Go has a dedicated concurrency model that makes it particularly useful for applications that require high performance,
scalability, and reliability.

### Rules

- If you don’t need it, don’t use it.
- Don’t communicate by sharing memory, share memory by communicating.
- Don’t over-engineer things by using shared memory and complicated, `error-prone` synchronization primitives,
  instead, use `message-passing` between go routines so variables and data can be used in the appropriate sequence.

### Points

- The fundamentals of concurrency in Go - Goroutines
- Techniques that can be used to write concurrent programs
- Concurrency patterns
- Advantages
- Pitfalls

### Sections

#### Synchronization: Mutexes & WaitGroups
- [Mutex](./mutexes_waitgroups/mutex)
- [WaitGroup](./mutexes_waitgroups/waitgroup)
- **Problems**:
  - [The Account Balance](./mutexes_waitgroups/account_balance)
  - [The Dining Philosophers](./mutexes_waitgroups/dining_philosophers)

#### Channels
- [Introduction](./channels)
- [Select Statement](./channels/select_statement)
- [Buffered Channel](./channels/buffered_channels)
- **Problems**:
  - [Producer–consumer problem](./channels/producer_consumer)
  - [The Sleeping Barber problem](./channels/sleeping_barber)
