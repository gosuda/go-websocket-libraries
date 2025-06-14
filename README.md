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
**Last Updated:** Sat, 14 Jun 2025 03:34:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.93 | 37.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.81 | 25.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.43 | 16.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 358.75 | 21.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 550.16 | 14.14 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)