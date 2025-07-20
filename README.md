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
**Last Updated:** Sun, 20 Jul 2025 03:57:35 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.81 | 37.40 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.83 | 25.87 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 468.10 | 16.53 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 364.51 | 21.24 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.80 | 13.91 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)