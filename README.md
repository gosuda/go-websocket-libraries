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
**Last Updated:** Sun, 08 Jun 2025 03:42:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 195.77 | 39.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.63 | 27.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 432.93 | 17.83 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 340.61 | 22.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 521.40 | 14.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)