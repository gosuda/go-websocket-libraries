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
**Last Updated:** Fri, 25 Sep 2026 08:10:44 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 247.45 | 31.23 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 267.18 | 29.02 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 487.71 | 15.89 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 336.75 | 22.99 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 563.40 | 13.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)