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
**Last Updated:** Sat, 28 Jun 2025 03:35:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 203.44 | 37.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.18 | 26.63 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 437.84 | 17.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 345.23 | 22.40 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 539.64 | 14.39 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)