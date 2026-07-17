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
**Last Updated:** Fri, 17 Jul 2026 05:29:11 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 257.65 | 30.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.75 | 27.35 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 505.50 | 15.31 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 345.11 | 22.40 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 590.84 | 13.16 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)