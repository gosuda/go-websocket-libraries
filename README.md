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
**Last Updated:** Fri, 13 Feb 2026 04:35:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.32 | 33.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.07 | 27.71 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.62 | 15.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 332.90 | 23.25 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 555.67 | 13.98 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)