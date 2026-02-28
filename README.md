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
**Last Updated:** Sat, 28 Feb 2026 04:08:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 234.39 | 33.06 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.81 | 27.45 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 518.11 | 14.94 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 354.70 | 21.82 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 580.54 | 13.39 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)