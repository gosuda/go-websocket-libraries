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
**Last Updated:** Thu, 22 Jan 2026 04:03:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 226.20 | 34.22 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.48 | 27.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 519.23 | 14.90 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 359.18 | 21.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 593.00 | 13.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)