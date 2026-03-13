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
**Last Updated:** Fri, 13 Mar 2026 04:26:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 236.43 | 32.74 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.73 | 27.33 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.63 | 14.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 361.08 | 21.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 590.01 | 13.17 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)