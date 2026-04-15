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
**Last Updated:** Wed, 15 Apr 2026 05:13:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.70 | 33.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.96 | 27.95 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 505.84 | 15.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 332.40 | 23.29 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 577.23 | 13.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)