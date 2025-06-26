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
**Last Updated:** Thu, 26 Jun 2025 03:39:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.96 | 37.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 306.70 | 25.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 462.80 | 16.71 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 361.81 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 555.15 | 13.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)