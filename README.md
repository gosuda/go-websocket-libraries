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
**Last Updated:** Sat, 16 May 2026 05:41:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 247.57 | 31.26 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.32 | 25.72 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 532.40 | 14.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 380.33 | 20.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 607.57 | 12.79 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)