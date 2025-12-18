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
**Last Updated:** Thu, 18 Dec 2025 09:11:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.87 | 34.25 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.82 | 28.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 517.74 | 14.94 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 343.92 | 22.50 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 575.50 | 13.52 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)