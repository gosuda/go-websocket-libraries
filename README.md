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
**Last Updated:** Mon, 02 Jun 2025 03:42:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.98 | 38.39 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.19 | 26.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 441.50 | 17.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 344.31 | 22.44 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.74 | 14.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)