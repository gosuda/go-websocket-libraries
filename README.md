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
**Last Updated:** Thu, 01 Jan 2026 04:03:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 267.71 | 28.96 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 376.67 | 20.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 626.82 | 12.35 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 444.31 | 17.46 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 788.79 | 9.86 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)