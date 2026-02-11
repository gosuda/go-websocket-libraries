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
**Last Updated:** Wed, 11 Feb 2026 04:54:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 251.63 | 30.78 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 313.26 | 24.71 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 571.55 | 13.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 390.22 | 19.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 656.59 | 11.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)