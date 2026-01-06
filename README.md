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
**Last Updated:** Tue, 06 Jan 2026 03:53:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.20 | 33.60 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.03 | 27.02 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 524.92 | 14.73 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 356.83 | 21.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 605.15 | 12.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)