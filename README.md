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
**Last Updated:** Fri, 24 Oct 2025 03:24:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.49 | 34.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.83 | 26.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 460.38 | 16.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 356.08 | 21.73 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.34 | 14.14 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)