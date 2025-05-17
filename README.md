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
**Last Updated:** Sat, 17 May 2025 03:29:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.31 | 37.65 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.21 | 26.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.47 | 16.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 360.71 | 21.43 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.44 | 14.01 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)