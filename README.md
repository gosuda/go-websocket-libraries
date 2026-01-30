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
**Last Updated:** Fri, 30 Jan 2026 04:23:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 249.73 | 30.97 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 309.07 | 25.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 547.41 | 14.14 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 394.74 | 19.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 631.26 | 12.30 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)