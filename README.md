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
**Last Updated:** Mon, 06 Jul 2026 06:54:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 265.36 | 29.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.16 | 26.21 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 532.82 | 14.53 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 367.90 | 21.06 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 609.02 | 12.76 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)