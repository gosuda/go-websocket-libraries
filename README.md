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
**Last Updated:** Wed, 11 Jun 2025 03:38:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.67 | 38.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.06 | 27.39 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 460.27 | 16.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 351.67 | 21.97 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 531.47 | 14.63 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)