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
**Last Updated:** Thu, 06 Aug 2026 05:35:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 258.90 | 29.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 269.80 | 28.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.18 | 15.85 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 329.94 | 23.46 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 574.61 | 13.55 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)