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
**Last Updated:** Fri, 03 Jul 2026 06:17:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 255.14 | 30.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 258.31 | 29.99 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 469.57 | 16.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 334.95 | 23.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 542.08 | 14.33 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)