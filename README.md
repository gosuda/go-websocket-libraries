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
**Last Updated:** Fri, 02 Jan 2026 03:54:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 251.04 | 30.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 310.53 | 24.93 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 556.03 | 13.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 390.14 | 19.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 629.65 | 12.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)