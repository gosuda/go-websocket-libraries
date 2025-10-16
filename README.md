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
**Last Updated:** Thu, 16 Oct 2025 03:27:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 213.17 | 36.13 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 269.76 | 28.65 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 421.38 | 18.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 330.59 | 23.41 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 505.13 | 15.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)