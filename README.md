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
**Last Updated:** Fri, 15 Aug 2025 03:43:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 194.75 | 39.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.50 | 27.41 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.58 | 17.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 346.83 | 22.31 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 541.32 | 14.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)