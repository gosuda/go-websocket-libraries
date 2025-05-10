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
**Last Updated:** Sat, 10 May 2025 03:27:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 194.95 | 39.64 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.01 | 27.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 425.25 | 18.19 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 331.16 | 23.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 510.15 | 15.20 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)