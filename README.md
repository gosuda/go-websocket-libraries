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
**Last Updated:** Fri, 07 Aug 2026 04:38:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 250.84 | 30.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 267.08 | 28.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 474.30 | 16.33 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 340.56 | 22.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 538.70 | 14.44 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)