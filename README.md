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
**Last Updated:** Tue, 13 May 2025 03:33:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.50 | 37.23 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.40 | 25.92 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 448.30 | 17.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 359.54 | 21.50 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 553.82 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)