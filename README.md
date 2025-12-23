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
**Last Updated:** Tue, 23 Dec 2025 03:51:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.95 | 33.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.38 | 27.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 529.78 | 14.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 355.38 | 21.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 589.20 | 13.20 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)