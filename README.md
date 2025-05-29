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
**Last Updated:** Thu, 29 May 2025 03:35:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 194.42 | 39.53 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.85 | 26.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 425.10 | 18.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 338.31 | 22.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 515.58 | 15.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)