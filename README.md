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
**Last Updated:** Fri, 13 Nov 2026 09:57:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 240.13 | 32.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 265.65 | 29.11 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 490.65 | 15.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 322.88 | 23.95 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 562.20 | 13.83 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)