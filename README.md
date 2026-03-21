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
**Last Updated:** Sat, 21 Mar 2026 04:19:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 220.51 | 35.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 264.53 | 29.22 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 473.51 | 16.31 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 325.84 | 23.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 539.13 | 14.42 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)