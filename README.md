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
**Last Updated:** Thu, 26 Mar 2026 04:57:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.46 | 34.01 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.40 | 27.94 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 485.04 | 15.96 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 343.48 | 22.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 549.08 | 14.13 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)