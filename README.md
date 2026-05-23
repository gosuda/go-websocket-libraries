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
**Last Updated:** Sat, 23 May 2026 05:55:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 236.79 | 32.67 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.17 | 26.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.04 | 14.83 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 357.31 | 21.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 606.94 | 12.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)