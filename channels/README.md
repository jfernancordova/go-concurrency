## Channels

Atomic functions and mutexes work, but they don’t make writing concurrent programs easier, less error-prone, or fun. 
In Go you don’t have only atomic functions and mutexes to keep shared resources safe and eliminate race conditions. You also have
channels that synchronize goroutines as they send and receive the resources they
need to share between each other.

When a resource needs to be shared between goroutines, channels act as a conduit
between the goroutines and provide a mechanism that guarantees a synchronous
exchange. When declaring a channel, the type of data that will be shared needs to be
specified. Values and pointers of built-in, named, struct, and reference types can be
shared through a channel.

### Bibliography

- [Go in Action](https://www.manning.com/books/go-in-action)
