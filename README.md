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
**Last Updated:** Fri, 15 May 2026 06:19:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.75 | 32.82 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.90 | 27.09 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 521.81 | 14.83 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 351.72 | 22.02 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 599.47 | 12.98 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)