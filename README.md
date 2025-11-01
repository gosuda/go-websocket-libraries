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
**Last Updated:** Sat, 01 Nov 2025 03:33:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.29 | 33.67 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.26 | 25.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 479.04 | 16.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 366.61 | 21.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 577.45 | 13.47 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)