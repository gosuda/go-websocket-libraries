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
**Last Updated:** Sat, 31 Oct 2026 08:50:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 441.05 | 17.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 482.75 | 16.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 844.29 | 9.20 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 567.24 | 13.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 1027.35 | 7.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)