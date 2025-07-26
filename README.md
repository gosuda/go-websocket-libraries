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
**Last Updated:** Sat, 26 Jul 2025 03:48:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 212.69 | 36.37 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 307.42 | 25.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 478.74 | 16.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 370.33 | 20.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 571.94 | 13.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)