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
**Last Updated:** Sun, 15 Jun 2025 03:42:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.60 | 37.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.91 | 25.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 458.15 | 16.85 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 355.19 | 21.77 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.70 | 14.11 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)