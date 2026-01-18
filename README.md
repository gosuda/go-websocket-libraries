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
**Last Updated:** Sun, 18 Jan 2026 03:59:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 229.79 | 33.67 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.72 | 27.09 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 525.59 | 14.74 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 362.70 | 21.35 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 600.76 | 12.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)