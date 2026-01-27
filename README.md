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
**Last Updated:** Tue, 27 Jan 2026 04:01:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 226.83 | 34.06 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.54 | 26.31 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 553.43 | 14.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 355.11 | 21.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 634.01 | 12.24 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)