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
**Last Updated:** Wed, 18 Mar 2026 04:36:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 232.79 | 33.23 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.89 | 27.83 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 501.09 | 15.42 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 349.07 | 22.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 553.43 | 14.05 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)