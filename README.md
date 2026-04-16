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
**Last Updated:** Thu, 16 Apr 2026 05:18:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.30 | 33.47 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 273.30 | 28.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 476.26 | 16.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 342.71 | 22.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 549.69 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)