# String data command
SET test1 "biji"
get test1
exists test1
del test1
append test1 " kuda"
keys test* -- take all keys that have prefix "test"

# range
getrange test3 0 2
setrange test3 0 bijia

# multiple
mget test1 test2
mset test7 aw test8 uw

# expiration
expire test7 8 -- existing key
setex test10 18 whatever
ttl key

# Increment and Decrement
```SHELL
incr key
decr key
incrby key [increment]
decrby key [decrement]
```

# Flush
flushdb -- cleanup only in our current db
flushall

# Pipeline

```
cat /usr/local/etc/command/sets.txt | redis-cli -h localhost --pipe
```

# transaction
multi
OK
localhost:6379(TX)> set tx1 "meow"
QUEUED
localhost:6379(TX)> set tx2 "mbew"
QUEUED
localhost:6379(TX)> exec
1) OK
2) OK

1 error

localhost:6379> multi
OK
localhost:6379(TX)> set tx4 "ssmmm"
QUEUED
localhost:6379(TX)> set txhh
localhost:6379(TX)> set txhh "sssl"
QUEUED
localhost:6379(TX)> exec
(error) EXECABORT Transaction discarded because of previous errors.