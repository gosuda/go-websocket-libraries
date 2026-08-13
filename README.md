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
**Last Updated:** Thu, 13 Aug 2026 04:29:06 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 267.81 | 28.90 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.75 | 26.63 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 507.83 | 15.28 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 367.95 | 21.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 584.53 | 13.28 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)