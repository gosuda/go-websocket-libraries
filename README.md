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
**Last Updated:** Fri, 27 Feb 2026 04:27:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 234.01 | 33.08 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.60 | 27.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 502.81 | 15.41 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 349.64 | 22.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 571.66 | 13.58 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)