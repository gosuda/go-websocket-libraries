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
**Last Updated:** Tue, 17 Feb 2026 04:32:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 222.60 | 34.72 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 266.06 | 29.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 491.32 | 15.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 328.93 | 23.50 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.88 | 13.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)