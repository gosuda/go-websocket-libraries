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
**Last Updated:** Thu, 27 Aug 2026 13:05:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 261.53 | 29.60 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.38 | 26.80 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 509.10 | 15.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 359.16 | 21.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 596.15 | 13.04 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)