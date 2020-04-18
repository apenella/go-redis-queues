Go Redis Queues
======

Go redis queues provide several queues solutions using Redis.
The purpose was explore and understand how redis works and how to implement a Golang application to interactuate with Redis.
The application uses [cobra](https://github.com/spf13/cobra) as cli and accept the commands listed below:

```
go-redis-queues interactuates with Redis and its types which could be used as queue.
 These redis types are:
 - List
 - Pubsub
 - Stream

Usage:
  go-redis-queues [flags]
  go-redis-queues [command]

Available Commands:
  consume     Consume message from Redis
  get         Provide information related to go-redis-queues items
  help        Help about any command
  publish     Publish message to Redis

Flags:
  -h, --help   help for go-redis-queues

Use "go-redis-queues [command] --help" for more information about a command.
```

When the commands **publish** or **consume** are executed, the kind of transport to use could be set using the flag *--transport*.
Take care to use the same transport for publisher and consumers.

```
Publish message to Redis

Usage:
  go-redis-queues publish [flags]

Flags:
  -c, --channel string     Channel where the message will be published to.
  -g, --group string       Consumer group used on stream transport (default "default")
  -h, --help               help for publish
  -m, --message string     Message to be published
  -t, --transport string   Choose transport type [pubsub(default)|stream|fifo]
```

## FIFO
- #### Publisher: 
Uses [*LPUSH*](https://redis.io/commands/lpush) to write messages to lists.
- #### Consumer:
Uses [*BRPOP*](https://redis.io/commands/brpop) to read messages from lists.

## PubSub
- #### Publisher:
Uses [*PUBLISH*](https://redis.io/commands/publish) to write messages to lists.
- #### Consumer:
Uses [*SUBSCRIBE*](https://redis.io/commands/subscribe) to read messages from lists.

## Streams
[Redis streams](https://redis.io/topics/streams-intro) used for these proposal are configured to with consumer groups. All streams uses a group named *default*.

- #### Publisher:
Uses [*XADD*](https://redis.io/commands/xadd) to write messages to lists.
- #### Consumer:
Uses [*XREADGROUP*](https://redis.io/commands/xreadgroup) to read messages from lists.