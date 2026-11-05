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
**Last Updated:** Thu, 05 Nov 2026 09:49:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 238.11 | 32.40 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 265.66 | 29.15 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 471.01 | 16.39 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 326.53 | 23.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 579.17 | 13.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)