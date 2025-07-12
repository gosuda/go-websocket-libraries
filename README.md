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
**Last Updated:** Sat, 12 Jul 2025 03:45:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.61 | 38.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.15 | 26.65 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 454.15 | 17.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 356.86 | 21.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.00 | 14.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)