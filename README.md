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
**Last Updated:** Thu, 11 Jun 2026 07:17:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.49 | 32.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.95 | 27.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.05 | 15.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 356.94 | 21.72 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 583.37 | 13.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)