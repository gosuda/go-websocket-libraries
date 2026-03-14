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
**Last Updated:** Sat, 14 Mar 2026 04:24:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.61 | 32.89 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.71 | 28.09 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 503.95 | 15.37 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 356.36 | 21.74 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 582.58 | 13.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)