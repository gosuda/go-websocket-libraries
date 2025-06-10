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
**Last Updated:** Tue, 10 Jun 2025 03:38:40 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 195.01 | 39.55 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.30 | 27.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 454.24 | 17.00 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 339.37 | 22.76 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.61 | 14.58 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)