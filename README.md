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
**Last Updated:** Mon, 30 Jun 2025 03:45:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 203.31 | 37.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.11 | 26.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 457.67 | 16.89 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 359.10 | 21.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.94 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)