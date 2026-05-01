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
**Last Updated:** Fri, 01 May 2026 05:56:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.21 | 32.01 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.91 | 26.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.81 | 14.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 354.23 | 21.87 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 610.98 | 12.73 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)