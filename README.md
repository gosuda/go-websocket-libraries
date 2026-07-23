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
**Last Updated:** Thu, 23 Jul 2026 05:41:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 317.09 | 24.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 399.66 | 19.41 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 610.18 | 12.71 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 441.62 | 17.55 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 784.06 | 9.92 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)