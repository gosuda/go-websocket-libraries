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
**Last Updated:** Thu, 10 Jul 2025 03:43:35 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 198.83 | 38.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.29 | 27.39 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 429.46 | 17.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 346.54 | 22.35 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 519.62 | 14.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)