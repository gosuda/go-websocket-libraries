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
**Last Updated:** Wed, 28 May 2025 03:34:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.84 | 38.40 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.92 | 26.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 457.68 | 16.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 352.90 | 21.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 534.10 | 14.50 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)