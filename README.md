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
**Last Updated:** Fri, 17 Apr 2026 05:17:06 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 268.23 | 28.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 367.24 | 21.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 596.07 | 13.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 411.11 | 18.85 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 719.23 | 10.80 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)