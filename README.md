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
**Last Updated:** Tue, 05 May 2026 05:31:40 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.27 | 33.62 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 273.72 | 28.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 499.57 | 15.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 354.04 | 21.85 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 570.30 | 13.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)