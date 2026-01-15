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
**Last Updated:** Thu, 15 Jan 2026 03:56:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 218.63 | 35.36 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 274.43 | 28.20 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 492.98 | 15.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 326.89 | 23.62 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 565.04 | 13.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)