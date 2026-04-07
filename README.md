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
**Last Updated:** Tue, 07 Apr 2026 05:00:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 228.27 | 33.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.20 | 27.32 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 497.53 | 15.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 330.69 | 23.42 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 561.26 | 13.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)