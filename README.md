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
**Last Updated:** Sun, 21 Jun 2026 07:19:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 264.61 | 29.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.65 | 25.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 546.24 | 14.19 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 367.73 | 21.05 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 626.43 | 12.39 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)