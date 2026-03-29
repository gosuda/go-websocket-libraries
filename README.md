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
**Last Updated:** Sun, 29 Mar 2026 05:03:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.57 | 34.05 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.05 | 27.22 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 501.52 | 15.42 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 346.57 | 22.34 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 600.72 | 12.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)