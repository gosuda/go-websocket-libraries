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
**Last Updated:** Thu, 20 Aug 2026 03:28:00 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 378.96 | 20.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 457.29 | 16.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 752.43 | 10.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 529.59 | 14.65 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 957.53 | 8.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)