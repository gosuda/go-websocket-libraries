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
**Last Updated:** Wed, 02 Jul 2025 03:41:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 199.85 | 38.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.70 | 26.94 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 439.41 | 17.59 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 343.93 | 22.51 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 525.50 | 14.79 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)