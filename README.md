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
**Last Updated:** Sat, 11 Apr 2026 04:36:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 247.89 | 31.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.20 | 26.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 537.17 | 14.42 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 358.57 | 21.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 623.06 | 12.49 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)