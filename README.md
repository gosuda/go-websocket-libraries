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
**Last Updated:** Sat, 18 Apr 2026 04:57:26 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 240.15 | 32.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.29 | 26.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 527.47 | 14.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 353.93 | 21.87 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 596.09 | 13.05 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)