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
**Last Updated:** Mon, 10 Aug 2026 04:16:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 268.09 | 28.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.58 | 26.10 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.76 | 15.00 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 358.93 | 21.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 601.10 | 12.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)