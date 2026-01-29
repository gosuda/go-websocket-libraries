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
**Last Updated:** Thu, 29 Jan 2026 04:20:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.66 | 31.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.74 | 27.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 532.43 | 14.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 370.83 | 20.85 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 615.38 | 12.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)