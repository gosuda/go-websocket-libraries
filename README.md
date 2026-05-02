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
**Last Updated:** Sat, 02 May 2026 05:25:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 233.58 | 33.07 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.98 | 26.72 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 508.92 | 15.20 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 339.53 | 22.77 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 572.26 | 13.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)