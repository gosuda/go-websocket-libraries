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
**Last Updated:** Thu, 16 Jul 2026 05:27:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 267.23 | 28.99 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.64 | 26.51 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 525.62 | 14.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 359.93 | 21.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 604.27 | 12.86 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)