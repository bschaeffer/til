# TCP Server

Wanted to learn how to build a TCP Server. Exact same as this blog post:
https://blog.oestrich.org/2017/07/using-ranch-with-elixir/

### Usage

Start the server

```bash
$ TCP_PORT=9876 mix run --no-halt
```

```bash
$ telnet localhost 9876
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Hi
Hi^]
telnet> Connection closed.
```
