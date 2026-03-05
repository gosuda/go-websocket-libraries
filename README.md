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
**Last Updated:** Thu, 05 Mar 2026 04:26:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.39 | 34.04 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.77 | 28.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 513.08 | 15.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 346.64 | 22.34 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 566.82 | 13.71 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)