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
**Last Updated:** Fri, 21 Nov 2025 03:34:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 217.75 | 35.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.60 | 27.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 441.03 | 17.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 343.79 | 22.51 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 531.24 | 14.63 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)