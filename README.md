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
**Last Updated:** Sat, 20 Jun 2026 06:48:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 264.96 | 29.23 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.71 | 26.54 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.97 | 15.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 367.49 | 21.07 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 597.40 | 13.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)