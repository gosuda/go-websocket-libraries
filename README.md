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
**Last Updated:** Tue, 08 Sep 2026 07:25:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.55 | 29.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.82 | 26.67 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 517.39 | 14.95 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 355.27 | 21.77 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 603.50 | 12.87 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)