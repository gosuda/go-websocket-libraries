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
**Last Updated:** Sun, 27 Jul 2025 03:58:48 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.86 | 37.29 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.05 | 25.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 455.77 | 16.96 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 358.25 | 21.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 552.29 | 14.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)