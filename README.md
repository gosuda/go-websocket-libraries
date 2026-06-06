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
**Last Updated:** Sat, 06 Jun 2026 06:10:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 234.15 | 33.03 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.99 | 27.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 498.65 | 15.52 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 332.85 | 23.24 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 564.14 | 13.78 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)