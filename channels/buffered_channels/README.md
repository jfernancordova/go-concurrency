## Buffered Channels

A buffered channel is a channel with capacity to hold one or more values before they’re
received. These types of channels don’t force goroutines to be ready at the same
instant to perform sends and receives. There are also different conditions for when a
send or receive does block. A receive will block only if there’s no value in the channel
to receive. A send will block only if there’s no available buffer to place the value being
sent. This leads to the one big difference between unbuffered and buffered channels:
An unbuffered channel provides a guarantee that an exchange between two gorou-
tines is performed at the instant the send and receive take place. A buffered channel
has no such guarantee.

## Unbuffered channels

An unbuffered channel is a channel with no capacity to hold any value before it’s
received. These types of channels require both a sending and receiving goroutine to
be ready at the same instant before any send or receive operation can complete. If the
two goroutines aren’t ready at the same instant, the channel makes the goroutine that
performs its respective send or receive operation first wait. Synchronization is inher-
ent in the interaction between the send and receive on the channel. One can’t hap-
pen without the other.

### Bibliography

- [Go in Action](https://www.manning.com/books/go-in-action)

