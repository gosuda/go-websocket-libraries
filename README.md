# go-websocket-libraries

Go WebSocket Library Comparison

**Libraries Benchmarked:**

- https://github.com/gorilla/websocket
- https://github.com/lxzan/gws
- https://github.com/gobwas/ws
- https://github.com/lesismal/nbio
- https://github.com/coder/websocket

**Latest Benchmark Results:**

<!-- BENCHMARK_TABLE_START -->
**Last Updated:** Fri, 09 May 2025 03:31:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.20 | 38.59 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.19 | 26.90 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 444.32 | 17.40 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 349.34 | 22.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 533.79 | 14.55 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)