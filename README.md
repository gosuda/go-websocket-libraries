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
**Last Updated:** Wed, 01 Apr 2026 05:14:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 232.07 | 33.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.24 | 27.24 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.72 | 14.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 339.41 | 22.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 593.21 | 13.08 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)