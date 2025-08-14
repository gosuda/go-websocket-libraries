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
**Last Updated:** Thu, 14 Aug 2025 03:42:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.00 | 37.75 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 299.63 | 25.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 475.57 | 16.28 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 361.41 | 21.42 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 571.86 | 13.60 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)