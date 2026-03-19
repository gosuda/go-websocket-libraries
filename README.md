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
**Last Updated:** Thu, 19 Mar 2026 04:36:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.46 | 32.04 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.23 | 26.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 512.09 | 15.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 356.59 | 21.73 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 590.70 | 13.18 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)