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
**Last Updated:** Wed, 29 Apr 2026 05:41:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.31 | 32.87 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 302.13 | 25.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 500.09 | 15.48 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 341.03 | 22.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 562.89 | 13.80 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)