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
**Last Updated:** Fri, 04 Jul 2025 03:39:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.53 | 37.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.16 | 26.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 460.41 | 16.78 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 359.86 | 21.49 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 552.52 | 14.08 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)