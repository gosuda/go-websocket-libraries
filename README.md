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
**Last Updated:** Wed, 04 Jun 2025 03:37:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.99 | 38.45 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.71 | 26.50 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 434.51 | 17.76 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 342.34 | 22.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 518.49 | 14.98 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)