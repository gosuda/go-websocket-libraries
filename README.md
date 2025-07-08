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
**Last Updated:** Tue, 08 Jul 2025 03:41:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 199.01 | 38.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.41 | 27.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.01 | 17.07 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 343.70 | 22.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.60 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)