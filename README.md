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
**Last Updated:** Wed, 24 Sep 2025 03:21:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.94 | 33.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 306.08 | 25.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 471.78 | 16.38 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 361.23 | 21.38 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 565.95 | 13.74 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)