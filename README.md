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
**Last Updated:** Tue, 14 Apr 2026 05:12:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 245.34 | 31.56 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.51 | 26.10 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 520.51 | 14.88 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 355.23 | 21.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 604.58 | 12.87 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)