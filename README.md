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
**Last Updated:** Fri, 27 Jun 2025 03:40:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 199.69 | 38.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.55 | 27.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 441.12 | 17.48 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 346.70 | 22.28 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 531.42 | 14.63 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)