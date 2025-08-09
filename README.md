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
**Last Updated:** Sat, 09 Aug 2025 03:41:44 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 194.76 | 39.64 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.63 | 26.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.72 | 17.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 341.80 | 22.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.88 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)