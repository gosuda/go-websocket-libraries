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
**Last Updated:** Mon, 27 Apr 2026 05:41:05 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 247.73 | 31.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 305.47 | 25.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 526.35 | 14.70 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 358.72 | 21.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 610.47 | 12.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)