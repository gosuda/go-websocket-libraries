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
**Last Updated:** Fri, 29 May 2026 06:41:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.91 | 32.82 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.89 | 27.64 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 478.41 | 16.20 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 335.61 | 23.07 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 569.08 | 13.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)