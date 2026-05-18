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
**Last Updated:** Mon, 18 May 2026 06:39:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 259.01 | 29.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 338.61 | 22.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 520.90 | 14.87 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 391.25 | 19.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 649.24 | 11.99 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)