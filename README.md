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
**Last Updated:** Fri, 26 Sep 2025 03:22:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 212.87 | 36.35 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 267.91 | 28.84 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 437.57 | 17.71 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 327.09 | 23.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 501.01 | 15.51 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)