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
**Last Updated:** Fri, 30 May 2025 03:33:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.78 | 38.08 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.93 | 27.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.40 | 17.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 355.49 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.71 | 14.55 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)