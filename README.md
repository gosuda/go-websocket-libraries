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
**Last Updated:** Sat, 29 Aug 2026 09:22:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 261.69 | 29.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.75 | 26.87 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 498.34 | 15.53 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 356.41 | 21.73 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 590.96 | 13.13 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)