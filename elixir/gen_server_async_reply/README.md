# GenServer: Async call/reply

Example of multiple `GenServer.call` blocks waiting on a single reply from the
server process.

### Usage

```elixir
{:ok, pid} = MultiTask.start_link()

for key <- 1..5, _i <- 1..10 do
  # Simulate 5 different tokens being requested by 50 different processes. If we
  # assume each token request takes 1 second, all calls should receive a reply
  # within ~1 second as well.
  spawn(fn -> MultiTask.get(pid, key) |> IO.inspect() end)
end
```

### Use Case

My use case was a single OAuth Token Cache potentially handling requests for
multiple tokens and having to initialize them (request a token) at the same time
for multiple clients. For instance:

1. 5 workers start.
2. Worker #1 makes a blocking request for token from the cache process
3. Cache attempts to read from ets, but it's empty, so we enter the cache server
   and request the token.
4. Worker #1 awaits reply.
5. Workers #2-5 also request a token, but they are blocked by #1's request.
6. Request finishes. Worker #1 reply.
7. Worker #2 reply.
8. Worker #3 reply.
9. ...

This appeared to be a feature at first. The process would block while requesting
a token, write to the cache, return the token, and all other blocked processes
would read from the cache.

This technically works for one process per token, but when multiple, differnt
tokens are being requested, not so much.

So we can use the call/reply logic to allow for multiple requests being handled
by a single cache process.
