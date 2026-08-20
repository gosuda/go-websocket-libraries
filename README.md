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
**Last Updated:** Thu, 20 Aug 2026 04:33:24 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 471.91 | 16.47 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 574.12 | 13.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 958.14 | 8.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 659.20 | 11.79 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1166.47 | 6.67 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)