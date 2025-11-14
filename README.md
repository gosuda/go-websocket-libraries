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
**Last Updated:** Fri, 14 Nov 2025 03:36:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.38 | 34.31 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 299.11 | 25.87 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.36 | 16.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 360.60 | 21.50 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 569.90 | 13.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)