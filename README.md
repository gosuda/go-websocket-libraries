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
**Last Updated:** Fri, 12 Jun 2026 07:07:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 249.20 | 31.06 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 309.81 | 24.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 524.04 | 14.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 371.05 | 20.90 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 617.37 | 12.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)