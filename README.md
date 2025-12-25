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
**Last Updated:** Thu, 25 Dec 2025 03:50:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.21 | 31.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 299.17 | 25.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 553.99 | 13.97 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 376.59 | 20.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 634.03 | 12.24 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)