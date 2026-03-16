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
**Last Updated:** Mon, 16 Mar 2026 05:05:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 237.93 | 32.55 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.48 | 27.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 526.52 | 14.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 358.86 | 21.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 607.66 | 12.79 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)