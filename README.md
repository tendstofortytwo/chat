# chat

## server

```
go run .
```

## client

in terminal 1:

```
touch chat.log
tail -f chat.log
```

in terminal 2:

```
nc server 3000 >> chat.log
```

type in terminal 2, read chat from terminal 1.

## license

MIT license; see LICENSE.md.
