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
**Last Updated:** Mon, 05 May 2025 03:33:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 188.38 | 40.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 267.70 | 28.73 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 409.74 | 18.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 319.22 | 24.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 502.23 | 15.47 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)