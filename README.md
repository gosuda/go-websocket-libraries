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
**Last Updated:** Sat, 09 May 2026 05:37:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.67 | 34.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 264.12 | 29.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 476.38 | 16.24 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 311.97 | 24.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 542.07 | 14.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)