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
**Last Updated:** Sat, 19 Jul 2025 03:43:56 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 208.75 | 37.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 319.71 | 24.21 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 474.90 | 16.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.5 | 367.75 | 21.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.41 | 13.89 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)