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
**Last Updated:** Sun, 17 Aug 2025 03:45:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 193.53 | 39.97 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.39 | 27.87 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 424.95 | 18.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 336.57 | 22.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 511.92 | 15.18 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)