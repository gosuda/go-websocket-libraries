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
**Last Updated:** Fri, 11 Sep 2026 07:29:05 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 248.73 | 31.08 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 268.21 | 28.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 493.10 | 15.67 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 332.56 | 23.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 547.87 | 14.18 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)