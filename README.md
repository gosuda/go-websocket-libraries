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
**Last Updated:** Thu, 19 Jun 2025 03:38:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.91 | 37.54 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 304.02 | 25.46 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.50 | 17.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 360.80 | 21.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.64 | 14.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)