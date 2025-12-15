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
**Last Updated:** Mon, 15 Dec 2025 03:56:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 226.50 | 34.13 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.45 | 26.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 464.36 | 16.67 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 358.57 | 21.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 561.79 | 13.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)