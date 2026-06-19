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
**Last Updated:** Fri, 19 Jun 2026 08:17:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 263.30 | 29.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.30 | 26.47 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 480.82 | 16.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 355.66 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 583.59 | 13.31 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)