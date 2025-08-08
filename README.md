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
**Last Updated:** Fri, 08 Aug 2025 03:57:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 209.29 | 36.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 302.29 | 25.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 465.14 | 16.63 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 363.87 | 21.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 558.70 | 13.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)