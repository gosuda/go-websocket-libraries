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
**Last Updated:** Sat, 05 Jul 2025 03:36:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 209.15 | 37.01 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.49 | 25.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 461.72 | 16.76 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 361.71 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 546.85 | 14.23 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)