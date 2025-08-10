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
**Last Updated:** Sun, 10 Aug 2025 03:54:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 210.07 | 36.89 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.10 | 25.68 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 463.38 | 16.71 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 358.33 | 21.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 554.99 | 14.00 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)