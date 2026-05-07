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
**Last Updated:** Thu, 07 May 2026 05:48:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.83 | 32.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.29 | 27.94 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 528.32 | 14.67 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 348.89 | 22.16 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 599.08 | 12.99 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)