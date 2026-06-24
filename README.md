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
**Last Updated:** Wed, 24 Jun 2026 06:30:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 274.44 | 28.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.44 | 25.99 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 540.11 | 14.36 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 370.14 | 20.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 631.77 | 12.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)