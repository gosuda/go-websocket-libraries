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
**Last Updated:** Wed, 03 Jun 2026 07:25:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 245.58 | 31.53 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.79 | 25.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 534.82 | 14.47 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 365.52 | 21.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 608.15 | 12.80 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)