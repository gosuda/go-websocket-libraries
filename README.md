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
**Last Updated:** Mon, 15 Jun 2026 08:40:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 266.08 | 29.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.01 | 26.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.81 | 15.82 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 347.93 | 22.21 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 585.76 | 13.26 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)