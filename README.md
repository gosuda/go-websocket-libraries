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
**Last Updated:** Thu, 18 Jun 2026 07:18:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 436.50 | 17.82 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 452.41 | 17.18 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 803.56 | 9.67 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 565.92 | 13.73 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 978.13 | 7.96 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)