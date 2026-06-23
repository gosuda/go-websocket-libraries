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
**Last Updated:** Tue, 23 Jun 2026 06:34:05 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 255.44 | 30.26 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 274.00 | 28.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.41 | 15.85 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 345.09 | 22.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 577.31 | 13.44 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)