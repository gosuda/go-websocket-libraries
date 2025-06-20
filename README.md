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
**Last Updated:** Fri, 20 Jun 2025 03:37:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 195.88 | 39.35 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.14 | 27.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 431.43 | 17.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 330.70 | 23.42 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 513.01 | 15.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)