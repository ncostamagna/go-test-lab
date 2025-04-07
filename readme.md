# Suite
We can split **unit test** and **integration test**

## Lifecycles
Every test has its lifecycle (setup, before test, after test, cleanup). We can set code in each steps

# Synctest
Use cases
- Timeouts
- Tickers
- Goroutines

```sh
GOEXPERIMENT=synctest go test ./synctest/example/...
```

## Bubble
insolate environment which controls all this things: timers, tickers and goroutines. It has a virtual clock on-site you dont wait the time to finish.
<br/>
If you run a sleep, you dont need to wait this expecitic duration