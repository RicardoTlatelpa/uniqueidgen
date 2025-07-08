# uniqueidgen
Unique ID generator in distributed systems

- auto_increment attribute does not work in a distributed environment because a single database server is not large enough and generating unique IDs across multiple databases with minimal delay is challenging.

- IDs should fit into 64-bit

- The system should be able to generate 10,000 IDs per second

 - IDs are ordered by date

 - IDs are numerical values only

## Approach
- Twitter's snowflake approach

Since in our requirements we have 64 bits to work with, we can make use of those 64 bits:

- Sign bit: 1 bit
- Timestamp: 41 bits
- Datacenter ID: 5 bits, 2^5 = 32 datacenters
- Machine ID: 5 bits
- Sequence Number: 12 bits

Datacenter and machine IDs are chose at the startup time.
