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
**Last Updated:** Sat, 20 Dec 2025 03:39:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 228.45 | 33.84 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.34 | 27.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 505.57 | 15.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 352.88 | 21.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 594.84 | 13.06 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)