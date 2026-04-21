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
**Last Updated:** Tue, 21 Apr 2026 05:16:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 216.79 | 35.65 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 261.79 | 29.55 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 468.16 | 16.52 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 301.62 | 25.62 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 536.44 | 14.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)